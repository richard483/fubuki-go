package main

import (
	"context"
	"fubuki-go/config"
	"fubuki-go/controller"
	"fubuki-go/model"
	repository "fubuki-go/repository/impl"
	"fubuki-go/router"
	service "fubuki-go/service/impl"
	"log/slog"
	"os"

	"github.com/redis/go-redis/v9"
	genai "google.golang.org/genai"

	"net/http"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	db := config.NewDbConnection()
	if err := db.AutoMigrate(&model.History{}); err != nil {
		slog.Error("#main - error on auto migrating DB", "error", err.Error())
	}

	opts, err := redis.ParseURL(config.EnvRedisURI())
	if err != nil {
		slog.Error("#main - error on parsing redis URI", "error", err.Error())
	}

	redisClient := redis.NewClient(opts)

	ctx := context.TODO()
	geminiClient, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.EnvGeminiApiKey(),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		slog.Error("#main - error on initiating Google genai client", "error", err.Error())
	}

	cacheRedisRepository := repository.NewCacheRepository(redisClient)
	geminiHistoryRepository := repository.NewHistoryRepository(db)
	geminiService := service.NewGeminiService(geminiClient, geminiHistoryRepository, cacheRedisRepository)
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

	slog.Info("#main - Swagger served on (http || https)://(host):" + config.EnvPort() + "/swagger/index.html")

	serverError := server.ListenAndServe()

	slog.Error("#main - Server error on listening and serving", "error", serverError)
}
