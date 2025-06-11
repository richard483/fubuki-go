package main

import (
	"context"
	"fubuki-go/config"
	"fubuki-go/controller"
	"fubuki-go/model"
	repository "fubuki-go/repository/impl"
	"fubuki-go/router"
	service "fubuki-go/service/impl"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/option"

	"net/http"
)

func main() {

	db := config.NewDbConnection()
	if err := db.AutoMigrate(&model.History{}); err != nil {
		log.Println("#ERROR " + err.Error())
	}

	opts, err := redis.ParseURL(config.EnvRedisURI())
	if err != nil {
		log.Println("#ERROR " + err.Error())
	}

	redisClient := redis.NewClient(opts)

	ctx := context.TODO()
	geminiClient, err := genai.NewClient(ctx, option.WithAPIKey(config.EnvGeminiApiKey()))

	if err != nil {
		log.Println("#ERROR " + err.Error())
	}

	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			log.Println("#ERROR " + err.Error())
		}
	}(geminiClient)

	geminiHistoryRepository := repository.NewHistoryRepository(db)
	geminiService := service.NewGeminiService(geminiClient, geminiHistoryRepository, redisClient)
	helloWorldService := service.NewHelloWorldService()
	geminiHistoryService := service.NewGeminiHistoryService(geminiHistoryRepository)
	ollamaService := service.NewOllamaService()

	helloWorldController := controller.NewHelloWorldController(helloWorldService)
	geminiController := controller.NewGeminiController(geminiService)
	geminiHistoryController := controller.NewHistoryController(geminiHistoryService)
	ollamaController := controller.NewOllamaController(ollamaService)

	route := router.New(helloWorldController, geminiController, geminiHistoryController, ollamaController)

	server := &http.Server{
		Addr:    ":" + config.EnvPort(),
		Handler: route,
	}

	log.Println("Swagger served on http://localhost:" + config.EnvPort() + "/swagger/index.html")

	serverError := server.ListenAndServe()

	log.Panicln(serverError)
}
