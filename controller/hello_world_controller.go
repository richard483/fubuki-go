package controller

import (
	"fubuki-go/service"
	"fubuki-go/service_interface"

	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
	service_interface.HelloWorldServiceInterface
}

func NewHelloWorldController(service *service.HelloWorldService) *HelloWorldController {
	return &HelloWorldController{HelloWorldServiceInterface: service}
}

func (ctr *HelloWorldController) HelloWorld(c *gin.Context) {
	ctr.HelloWorldServiceInterface.HelloWorld(c)
}
