package controller

import (
	"fubuki-go/service"
	"fubuki-go/service_interface"
	"github.com/gin-gonic/gin"
)

type GeminiController struct {
	service_interface.GeminiServiceInterface
}

func NewGeminiController(service *service.GeminiService) *GeminiController {
	return &GeminiController{GeminiServiceInterface: service}
}

func (ctr *GeminiController) PromptText(c *gin.Context) {
	ctr.GeminiServiceInterface.PromptText(c)
	return
}

func (ctr *GeminiController) Chat(c *gin.Context) {
	ctr.GeminiServiceInterface.Chat(c)
	return
}
