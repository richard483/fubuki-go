package router

import (
	"fubuki-go/config"
	"fubuki-go/controller"
	"github.com/gin-gonic/gin"
)

func New(helloWorldController *controller.HelloWorldController, geminiController *controller.GeminiController, geminiHistoryController *controller.GeminiHistoryController) *gin.Engine {

	router := gin.Default()

	config.InitializeSwagger(router)

	misc := router.Group("/")
	{
		misc.GET("/hello-world", helloWorldController.HelloWorld)

	}

	gemini := router.Group("/gemini")
	{
		gemini.POST("/prompt-text", geminiController.PromptText)
		gemini.POST("/chat", geminiController.Chat)
		gemini.GET("/tune", geminiController.TuneModel)
	}

	geminiHistory := router.Group("/gemini-history")
	{
		geminiHistory.POST("/history-data", geminiHistoryController.CreateHistoryData)
		geminiHistory.PATCH("/history-data", geminiHistoryController.UpdateHistoryData)
		geminiHistory.GET("/history-data", geminiHistoryController.GetAllHistoryData)
		geminiHistory.DELETE("/history-data", geminiHistoryController.DeleteHistoryData)
	}

	return router
}
