package router

import (
	"fubuki-go/controller"
	"github.com/gin-gonic/gin"
)

func New(helloWorldController *controller.HelloWorldController) *gin.Engine {

	router := gin.Default()

	misc := router.Group("/")
	{
		misc.GET("/hello-world", helloWorldController.HelloWorld)
	}

	return router
}
