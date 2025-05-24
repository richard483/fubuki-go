package controller

import (
	"fubuki-go/service"

	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
	service.HelloWorldServiceInterface
}

func NewHelloWorldController(service service.HelloWorldServiceInterface) *HelloWorldController {
	return &HelloWorldController{HelloWorldServiceInterface: service}
}

func (ctr *HelloWorldController) HelloWorld(c *gin.Context) {
	ctr.HelloWorldServiceInterface.HelloWorld(c)
}
