package controller

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response"
	"fubuki-go/service"
	"fubuki-go/service_interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GeminiHistoryController struct {
	Service service_interface.GeminiHistoryServiceInterface
}

func NewGeminiHistoryController(service *service.GeminiHistoryService) *GeminiHistoryController {
	return &GeminiHistoryController{Service: service}
}

// CreateHistoryData godoc
// @Summary      CreateHistoryData
// @Description  create history data
// @Tags         gemini-history
// @Consume      json
// @Produce      json
// @Param        CreateGeminiHistory body request.CreateGeminiHistory true "Request Body"
// @Router       /gemini-history/history-data [post]
func (ctr *GeminiHistoryController) CreateHistoryData(c *gin.Context) {
	var historyData request.GeminiHistory
	if err := c.Bind(&historyData); err != nil {
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err := ctr.Service.CreateHistoryData(&historyData)

	if err != nil {
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
		Data:       "Success created history data",
	}
	c.JSON(http.StatusOK, res)
	return
}

// GetAllHistoryData godoc
// @Summary      GetAllHistoryData
// @Description  get all history data
// @Tags         gemini-history
// @Consume      json
// @Produce      json
// @Router       /gemini-history/history-data [get]
func (ctr *GeminiHistoryController) GetAllHistoryData(c *gin.Context) {
	data := ctr.Service.GetAllHistoryData()
	res := response.DefaultResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       data,
	}
	c.JSON(http.StatusOK, res)
	return
}

// UpdateHistoryData godoc
// @Summary      UpdateHistoryData
// @Description  update history data
// @Tags         gemini-history
// @Consume      json
// @Produce      json
// @Param        UpdateGeminiHistory body request.UpdateGeminiHistory true "Request Body"
// @Router       /gemini-history/history-data [patch]
func (ctr *GeminiHistoryController) UpdateHistoryData(c *gin.Context) {
	var historyData request.GeminiHistory
	if err := c.Bind(&historyData); err != nil {
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err := ctr.Service.UpdateHistoryData(&historyData)
	if err != nil {
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
		Data:       "Success updated history data",
	}
	c.JSON(http.StatusOK, res)
	return
}

// DeleteHistoryData godoc
// @Summary      DeleteHistoryData
// @Description  delete history data
// @Tags         gemini-history
// @Consume      json
// @Produce      json
// @Param        id query string false "history ID to be deleted"
// @Router       /gemini-history/history-data [delete]
func (ctr *GeminiHistoryController) DeleteHistoryData(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      "no 'id' parameter found on REST request",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := ctr.Service.DeleteHistoryData(id)

	if err != nil {
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
		Data:       "Success updated history data",
	}
	c.JSON(http.StatusOK, res)
	return
}
