package controller

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response"
	"fubuki-go/service"
	"log"
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
		log.Println(err)
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	data, err := ctr.OllamaServiceInterface.PromptOllamaText(&prompt)

	if err != nil {
		log.Println(err)
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
		log.Println(err)
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	data, err := ctr.OllamaServiceInterface.ChatOllama(&prompt)

	if err != nil {
		log.Println(err)
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
	err := ctr.OllamaServiceInterface.ResetChat()

	if err != nil {
		log.Println(err)
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
