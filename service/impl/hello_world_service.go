package impl

import (
	"github.com/gin-gonic/gin"
)

type HelloWorldService struct {
}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

func (srv *HelloWorldService) HelloWorld(c *gin.Context) {
	c.JSON(200, "Hello World!!")
}
