package service_interface

import (
	"github.com/gin-gonic/gin"
)

type GeminiServiceInterface interface {
	PromptText(c *gin.Context)
	Chat(c *gin.Context)
}
