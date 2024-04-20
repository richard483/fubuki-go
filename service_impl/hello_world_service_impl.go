package service_impl

import (
	"github.com/gin-gonic/gin"
)

type HelloWorldServiceImpl struct {
}

func (h HelloWorldServiceImpl) HelloWorld(c *gin.Context) {
	c.JSON(200, "Hello World!!")
	return
}
