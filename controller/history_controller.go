package controller

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response"
	"fubuki-go/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HistoryController struct {
	service.HistoryServiceInterface
}

func NewHistoryController(service service.HistoryServiceInterface) *HistoryController {
	return &HistoryController{HistoryServiceInterface: service}
}

// CreateHistoryData godoc
// @Summary      CreateHistoryData
// @Description  create history data
// @Tags         history
// @Consume      json
// @Produce      json
// @Param        CreateHistory body request.History true "Request Body"
// @Router       /history/data [post]
func (ctr *HistoryController) CreateHistoryData(c *gin.Context) {
	var historyData request.History
	if err := c.Bind(&historyData); err != nil {
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	slog.Info("#CreateHistoryData - processing request to create history data", "body", historyData, "path", c.Request.URL.Path)

	err := ctr.HistoryServiceInterface.CreateHistoryData(&historyData)

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
}

// CreateManyHistoryData godoc
// @Summary      CreateManyHistoryData
// @Description  create many history data
// @Tags         history
// @Consume      json
// @Produce      json
// @Param        CreateManyHistory body []request.History true "Request Body"
// @Router       /history/data/bulk [post]
func (ctr *HistoryController) CreateManyHistoryData(c *gin.Context) {
	var historyData []request.History
	if err := c.Bind(&historyData); err != nil {
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	slog.Info("#CreateManyHistoryData - processing request to create many history data", "body", historyData, "path", c.Request.URL.Path)

	err := ctr.HistoryServiceInterface.CreateManyHistoryData(&historyData)

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
		Data:       "Success created bulk history data",
	}
	c.JSON(http.StatusOK, res)
}

// GetAllHistoryData godoc
// @Summary      GetAllHistoryData
// @Description  get all history data by model source
// @Tags         history
// @Consume      json
// @Produce      json
// @Param        modelSource path string true "Model Source"
// @Router       /history/data/{modelSource} [get]
func (ctr *HistoryController) GetAllHistoryDataByModelSource(c *gin.Context) {
	slog.Info("#GetAllHistoryDataByModelSource - processing request to get all history data by model source", "modelSource", c.Param("modelSource"), "path", c.Request.URL.Path)
	modelSource := c.Param("modelSource")
	data := ctr.HistoryServiceInterface.GetAllHistoryDataByModelSource(modelSource)
	res := response.DefaultResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       data,
	}
	c.JSON(http.StatusOK, res)
}

// UpdateHistoryData godoc
// @Summary      UpdateHistoryData
// @Description  update history data
// @Tags         history
// @Consume      json
// @Produce      json
// @Param        UpdateHistory body request.UpdateHistory true "Request Body"
// @Router       /history/data [patch]
func (ctr *HistoryController) UpdateHistoryData(c *gin.Context) {
	var historyData request.UpdateHistory
	if err := c.Bind(&historyData); err != nil {
		res := response.DefaultResponse{
			StatusCode: http.StatusBadRequest,
			Message:    http.StatusText(http.StatusBadRequest),
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	slog.Info("#UpdateHistoryData - processing request to update history data", "body", historyData, "path", c.Request.URL.Path)
	err := ctr.HistoryServiceInterface.UpdateHistoryData(&historyData)
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
}

// DeleteHistoryData godoc
// @Summary      DeleteHistoryData
// @Description  delete history data
// @Tags         history
// @Consume      json
// @Produce      json
// @Param        id query string false "history ID to be deleted"
// @Router       /history/data [delete]
func (ctr *HistoryController) DeleteHistoryData(c *gin.Context) {
	slog.Info("#DeleteHistoryData - processing request to delete history data", "id", c.Query("id"), "path", c.Request.URL.Path)

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

	err := ctr.HistoryServiceInterface.DeleteHistoryData(id)

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
}
