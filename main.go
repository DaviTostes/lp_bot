package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/create", createLandingPage)

	router.Run()
}

type CreateRequest struct {
	CompanyInfo string `json:"company_info"`
}

func createLandingPage(ctx *gin.Context) {
	var reqBody CreateRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid body",
		})
		return
	}

	promptTemplate, err := os.ReadFile("prompt.txt")
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "couldn't read prompt file",
		})
		return
	}

	prompt := fmt.Sprintf(string(promptTemplate), reqBody.CompanyInfo)

	cmd := exec.Command("claude", "-p", "--model", "claude-opus-4-6", string(prompt))
	output, err := cmd.Output()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error on claude: " + err.Error(),
		})
		return
	}

	err = os.WriteFile("template/app/page.tsx", trimBullshit(output), 0644)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error writing file: " + err.Error(),
		})
		return
	}

	url, err := deployVercel()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error deploying on vercel: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"url":     url,
	})
}

func deployVercel() (string, error) {
	cmd := exec.Command("npm", "install")
	cmd.Dir = "./template"
	if err := cmd.Run(); err != nil {
		return "", err
	}

	cmd = exec.Command("vercel", "link", "--yes")
	cmd.Dir = "./template"
	if err := cmd.Run(); err != nil {
		return "", err
	}

	cmd = exec.Command("vercel", "--prod")
	cmd.Dir = "./template"
	url, err := cmd.Output()
	return string(url), err
}

func trimBullshit(output []byte) []byte {
	str := string(output)
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "```js")
	str = strings.TrimPrefix(str, "```javascript")
	str = strings.TrimPrefix(str, "```tsx")
	str = strings.TrimPrefix(str, "```typescript")
	str = strings.TrimSuffix(str, "```")
	str = strings.TrimSpace(str)
	return []byte(str)
}
