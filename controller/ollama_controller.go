package controller

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response"
	"fubuki-go/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OllamaController struct {
	service.OllamaServiceInterface
}

func NewOllamaController(service service.OllamaServiceInterface) *OllamaController {
	return &OllamaController{OllamaServiceInterface: service}
}

// PromptOllamaText godoc
// @Summary      Prompt Ollama Text
// @Description  get Ollama prompt text result by defining the model and text
// @Tags         ollama
// @Consume      json
// @Produce      json
// @Param        PromptText body request.PromptText true "Request Body"
// @Router       /ollama/prompt-text [post]
func (ctr *OllamaController) PromptOllamaText(c *gin.Context) {
	var prompt request.PromptText
	if err := c.Bind(&prompt); err != nil {
		slog.Error("#PromptOllamaText - error binding request", "error", err.Error())
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	slog.Info("#PromptOllamaText - processing request to get prompt text", "body", prompt, "path", c.Request.URL.Path)

	data, err := ctr.OllamaServiceInterface.PromptOllamaText(&prompt, c.Request.Context())

	if err != nil {
		slog.Error("#PromptOllamaText - error getting prompt text", "error", err.Error())
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
		Data:       &response.PromptTextData{Text: (*data).Response},
	}
	c.JSON(http.StatusOK, res)
}

// ChatOllama godoc
// @Summary      Chat with Ollama Model
// @Description  get Ollama chat result by defining the model and message
// @Tags         ollama
// @Consume      json
// @Produce      json
// @Param        PromptText body request.PromptText true "Request Body"
// @Router       /ollama/chat [post]
func (ctr *OllamaController) ChatOllama(c *gin.Context) {
	var prompt request.PromptText
	if err := c.Bind(&prompt); err != nil {
		slog.Error("#ChatOllama - error binding request", "error", err.Error())
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	slog.Info("#ChatOllama - processing request to get chat response", "body", prompt, "path", c.Request.URL.Path)

	data, err := ctr.OllamaServiceInterface.ChatOllama(&prompt, c.Request.Context())

	if err != nil {
		slog.Error("#ChatOllama - error getting chat response", "error", err.Error())
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
		Data:       &response.PromptTextData{Text: (*data).Message.Content},
	}
	c.JSON(http.StatusOK, res)
}

// ResetChat godoc
// @Summary      Reset Chat with Ollama Model
// @Description  reset chat with Ollama model
// @Tags         ollama
// @Consume      json
// @Produce      json
// @Router       /ollama/reset [get]
func (ctr *OllamaController) ResetChat(c *gin.Context) {
	slog.Info("#ResetChat - processing request to reset chat", "path", c.Request.URL.Path)
	err := ctr.OllamaServiceInterface.ResetChat()

	if err != nil {
		slog.Error("#ResetChat - error resetting chat", "error", err.Error())
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
		Data:       "Ollama chat reset successfully",
	}
	c.JSON(http.StatusOK, res)
}
