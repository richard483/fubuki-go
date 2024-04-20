package service_interface

import "github.com/gin-gonic/gin"

type HelloWorldServiceInterface interface {
	HelloWorld(c *gin.Context)
}
