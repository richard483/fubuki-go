package controller

import (
	"fubuki-go/service"
	"fubuki-go/service_interface"
	"github.com/gin-gonic/gin"
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
	ctr.Service.CreateHistoryData(c)
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
	ctr.Service.GetAllHistoryData(c)
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
	ctr.Service.UpdateHistoryData(c)
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
	ctr.Service.DeleteHistoryData(c)
	return
}
