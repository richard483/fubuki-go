package router

import (
	"fubuki-go/config"
	"fubuki-go/controller"

	"github.com/gin-gonic/gin"
)

func New(helloWorldController *controller.HelloWorldController, geminiController *controller.GeminiController, geminiHistoryController *controller.GeminiHistoryController, ollamaController *controller.OllamaController) *gin.Engine {

	if config.EnvReleaseMode() {
		gin.SetMode(gin.ReleaseMode)
	}

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
		gemini.GET("/reset", geminiController.ResetSession)
	}

	ollama := router.Group("/ollama")
	{
		ollama.POST("/prompt-text", ollamaController.PromptOllamaText)
	}

	geminiHistory := router.Group("/gemini-history")
	{
		geminiHistory.POST("/history-data", geminiHistoryController.CreateHistoryData)
		geminiHistory.POST("/history-data/bulk", geminiHistoryController.CreateManyHistoryData)
		geminiHistory.PATCH("/history-data", geminiHistoryController.UpdateHistoryData)
		geminiHistory.GET("/history-data", geminiHistoryController.GetAllHistoryData)
		geminiHistory.DELETE("/history-data", geminiHistoryController.DeleteHistoryData)
	}

	return router
}
