package router

import (
	"fubuki-go/controller"
	"github.com/gin-gonic/gin"
)

func New(helloWorldController *controller.HelloWorldController, geminiController *controller.GeminiController) *gin.Engine {

	router := gin.Default()

	misc := router.Group("/")
	{
		misc.GET("/hello-world", helloWorldController.HelloWorld)

	}

	gemini := router.Group("/gemini")
	{
		gemini.POST("/prompt-text", geminiController.PromptText)
		gemini.POST("/chat", geminiController.Chat)
	}

	return router
}
