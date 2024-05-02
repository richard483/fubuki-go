package service_interface

import (
	"github.com/gin-gonic/gin"
)

type GeminiHistoryServiceInterface interface {
	CreateHistoryData(c *gin.Context)
	GetAllHistoryData(c *gin.Context)
	UpdateHistoryData(c *gin.Context)
	DeleteHistoryData(c *gin.Context)
}
