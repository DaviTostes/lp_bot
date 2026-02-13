package handlers

import (
	"github.com/gin-gonic/gin"
	"lp_bot/internal/services"
)

type LPRequest struct {
	CompanyInfo string `json:"company_info"`
}

type LPResponse struct {
	Message string  `json:"message"`
	Error   *string `json:"error,omitempty"`
	Url     *string `json:"url,omitempty"`
}

func CreateLandingPage(ctx *gin.Context) {
	var body LPRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		errMsg := err.Error()
		ctx.JSON(400, LPResponse{
			Message: "Invalid Request",
			Error:   &errMsg,
		})
		return
	}

	url, err := services.CreateDeployLP(body.CompanyInfo)
	if err != nil {
		errMsg := err.Error()
		ctx.JSON(500, LPResponse{
			Message: "Error creating landing page",
			Error:   &errMsg,
		})
		return
	}

	ctx.JSON(201, LPResponse{
		Message: "Landing Page successfully deployed",
		Url:     &url,
	})
}

type CompanyInfoRequest struct {
	CompanyUrl string `json:"company_url"`
}

type CompanyInfoResponse struct {
	Message     string  `json:"message"`
	CompanyInfo *string `json:"company_info,omitempty"`
	Error       *string `json:"error,omitempty"`
}

func CreateCompanyInfo(ctx *gin.Context) {
	var body CompanyInfoRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		errMsg := err.Error()
		ctx.JSON(400, LPResponse{
			Message: "Invalid Request",
			Error:   &errMsg,
		})
		return
	}

	companyInfo, err := services.CreateCompanyInfo(ctx, body.CompanyUrl)
	if err != nil {
		errMsg := err.Error()
		ctx.JSON(400, LPResponse{
			Message: "Company info not generated",
			Error:   &errMsg,
		})
		return
	}

	ctx.JSON(201, CompanyInfoResponse{
		Message:     "Company info created!",
		CompanyInfo: &companyInfo,
	})
}
