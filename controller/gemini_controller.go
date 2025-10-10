package controller

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response"
	"fubuki-go/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeminiController struct {
	service.GeminiServiceInterface
}

func NewGeminiController(service service.GeminiServiceInterface) *GeminiController {
	return &GeminiController{GeminiServiceInterface: service}
}

// PromptText godoc
// @Summary      Prompt Text
// @Description  get prompt text result by prompt string
// @Tags         gemini
// @Consume      json
// @Produce      json
// @Param        PromptText body request.PromptText true "Request Body"
// @Router       /gemini/prompt-text [post]
func (ctr *GeminiController) PromptText(c *gin.Context) {
	var prompt request.PromptText
	if err := c.Bind(&prompt); err != nil {
		slog.Error("#PromptText - error binding request", "error", err.Error())
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	slog.Info("#PromptText - processing request to get prompt text", "body", prompt, "path", c.Request.URL.Path)

	data, err := ctr.GeminiServiceInterface.PromptText(&prompt, c.Request.Context())

	if err != nil {
		slog.Error("#PromptText - error getting prompt text", "error", err.Error())
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := response.DefaultResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       &response.PromptTextData{Text: (data)},
	}
	c.JSON(http.StatusOK, res)
}

// Chat godoc
// @Summary      Chat
// @Description  chat action API
// @Tags         gemini
// @Consume      json
// @Produce      json
// @Param        PromptText body request.PromptText true "Request Body"
// @Router       /gemini/chat [post]
func (ctr *GeminiController) Chat(c *gin.Context) {
	var prompt request.PromptText
	if err := c.Bind(&prompt); err != nil {
		slog.Error("#Chat - error binding request", "error", err.Error())
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	slog.Info("#Chat - processing request to get chat response", "body", prompt, "path", c.Request.URL.Path)

	data, err := ctr.GeminiServiceInterface.Chat(&prompt, c.Request.Context())
	if err != nil {
		slog.Error("#Chat - error getting chat response", "error", err.Error())
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := response.DefaultResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       &response.PromptTextData{Text: (data)},
	}
	c.JSON(http.StatusOK, res)
}

// ResetSession godoc
// @Summary      Reset chat session
// @Description  for resetting all chat session
// @Tags         gemini
// @Consume      json
// @Produce      json
// @Router       /gemini/reset [get]
func (ctr *GeminiController) ResetSession(c *gin.Context) {

	slog.Info("#ResetSession - processing request to reset chat session", "path", c.Request.URL.Path)

	data, err := ctr.GeminiServiceInterface.ResetSession(c.Request.Context())

	if err != nil {
		slog.Error("#ResetSession - error resetting session", "error", err.Error())
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.DefaultResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       &response.PromptTextData{Text: data},
	}
	c.JSON(http.StatusOK, res)
}
