package config

import (
	"fubuki-go/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "fubuki-go API"
	docs.SwaggerInfo.Description = "This is API for fubuki AI persona"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
