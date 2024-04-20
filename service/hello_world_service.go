package service

import "github.com/gin-gonic/gin"

type HelloWorldService interface {
	HelloWorld(c *gin.Context)
}
