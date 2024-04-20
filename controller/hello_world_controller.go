package controller

import (
	"fubuki-go/service"
	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
	service.HelloWorldService
}

func (ctr HelloWorldController) HelloWorld(c *gin.Context) {
	ctr.HelloWorldService.HelloWorld(c)
	return
}
