package service

import "github.com/gin-gonic/gin"

type HelloWorldServiceInterface interface {
	HelloWorld(c *gin.Context)
}
