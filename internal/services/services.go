package services

import (
	"lp_bot/internal/gen"
	"lp_bot/internal/vercel"
	"lp_bot/pkg/utils"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func CreateDeployLP(companyInfo string) (string, error) {
	output, err := gen.GenLP(companyInfo)
	if err != nil {
		return "", err
	}

	err = os.WriteFile("template/app/page.tsx", utils.TrimBullshit(output), 0644)
	if err != nil {
		return "", err
	}

	url, err := vercel.Deploy("./template")
	if err != nil {
		return "", err
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
