package router

import (
	"fubuki-go/config"
	"fubuki-go/controller"

	"github.com/gin-gonic/gin"
)

func New(helloWorldController *controller.HelloWorldController, geminiController *controller.GeminiController, geminiHistoryController *controller.HistoryController, ollamaController *controller.OllamaController) *gin.Engine {

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
		ollama.POST("/chat", ollamaController.ChatOllama)
		ollama.GET("/reset", ollamaController.ResetChat)
	}

	geminiHistory := router.Group("/history")
	{
		geminiHistory.POST("/data", geminiHistoryController.CreateHistoryData)
		geminiHistory.POST("/data/bulk", geminiHistoryController.CreateManyHistoryData)
		geminiHistory.PATCH("/data", geminiHistoryController.UpdateHistoryData)
		geminiHistory.GET("/data/:modelSource", geminiHistoryController.GetAllHistoryDataByModelSource)
		geminiHistory.DELETE("/data", geminiHistoryController.DeleteHistoryData)
	}

	return router
}
