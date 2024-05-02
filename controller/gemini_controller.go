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

// PromptText godoc
// @Summary      Prompt Text
// @Description  get prompt text result by prompt string
// @Tags         gemini
// @Consume      json
// @Produce      json
// @Param        GeminiText body request.GeminiText true "Request Body"
// @Router       /gemini/prompt-text [post]
func (ctr *GeminiController) PromptText(c *gin.Context) {
	ctr.GeminiServiceInterface.PromptText(c)
	return
}

// Chat godoc
// @Summary      Chat
// @Description  chat action API
// @Tags         gemini
// @Consume      json
// @Produce      json
// @Param        GeminiText body request.GeminiText true "Request Body"
// @Router       /gemini/chat [post]
func (ctr *GeminiController) Chat(c *gin.Context) {
	ctr.GeminiServiceInterface.Chat(c)
	return
}
