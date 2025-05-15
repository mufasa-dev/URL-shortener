package router

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitializeHandler()

	router.GET("/shorten", handlers.ShortenUrl)
	router.GET("/", handlers.RedirectHandler)

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
