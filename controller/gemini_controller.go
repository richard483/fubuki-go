package controller

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response"
	"fubuki-go/service_interface"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GeminiController struct {
	service_interface.GeminiServiceInterface
}

func NewGeminiController(service service_interface.GeminiServiceInterface) *GeminiController {
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
	var prompt request.GeminiText
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
	err, data := ctr.GeminiServiceInterface.PromptText(&prompt)

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
		Data:       &response.GeminiTextData{Text: (*data)[0]},
	}
	c.JSON(http.StatusOK, res)
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
	var prompt request.GeminiText
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

	err, data := ctr.GeminiServiceInterface.Chat(&prompt)
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
		Data:       &response.GeminiTextData{Text: (*data)[0]},
	}
	c.JSON(http.StatusOK, res)
	return
}

// ResetSession godoc
// @Summary      Reset chat session
// @Description  for resetting all chat session
// @Tags         gemini
// @Consume      json
// @Produce      json
// @Router       /gemini/reset [get]
func (ctr *GeminiController) ResetSession(c *gin.Context) {
	err, data := ctr.GeminiServiceInterface.ResetSession()

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
		Data:       &response.GeminiTextData{Text: data},
	}
	c.JSON(http.StatusOK, res)
	return
}
