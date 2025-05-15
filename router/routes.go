package router

import (
	"url-shortener/handlers"

	docs "url-shortener/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	handlers.InitializeHandler()

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/shorten", handlers.ShortenUrl)
	router.GET("/", handlers.RedirectHandler)

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
