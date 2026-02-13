package main

import (
	"lp_bot/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/create", handlers.CreateLandingPage)
	router.POST("/company", handlers.CreateCompanyInfo)

	router.Run()
}
