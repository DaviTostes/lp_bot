package services

import (
	"fmt"
	"lp_bot/internal/gen"
	"lp_bot/internal/vercel"
	"lp_bot/pkg/utils"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func CreateDeployLP(companyInfo string) (string, error) {
	templateDir := "./template"

	err := os.WriteFile(templateDir+"/app/page.tsx", []byte(""), 0744)
	if err != nil {
		return "", fmt.Errorf("Error clearing page.tsx file: %s", err)
	}

	if _, err := gen.GenLP(companyInfo, templateDir); err != nil {
		return "", fmt.Errorf("Error generating lp content: %s", err)
	}

	runCmd := exec.Command("npm", "run", "build")
	runCmd.Dir = templateDir
	if err := runCmd.Run(); err != nil {
		return "", fmt.Errorf("Error trying to build lp: %s", err)
	}

	url, err := vercel.Deploy(templateDir)
	if err != nil {
		return "", fmt.Errorf("Error trying to deploy lp: %s", err)
	}

	return url, nil
}

func CreateCompanyInfo(ctx *gin.Context, companyUrl string) (string, error) {
	cmd := exec.Command("curl", companyUrl)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	extractedInfo := utils.ExtractTextContent(string(output))

	companyInfo, err := gen.GenCompanyInfo(ctx, extractedInfo)
	return companyInfo, err
}
