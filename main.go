package main

import (
	"context"
	"fubuki-go/config"
	"fubuki-go/controller"
	"fubuki-go/model"
	"fubuki-go/repository"
	"fubuki-go/router"
	"fubuki-go/service"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"

	"net/http"
)

func main() {

	db := config.NewDbConnection()
	if err := db.AutoMigrate(&model.History{}); err != nil {
		log.Println("#ERROR " + err.Error())
	}

	ctx := context.TODO()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.EnvGeminiApiKey()))

	if err != nil {
		log.Println("#ERROR " + err.Error())
	}

	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			log.Println("#ERROR " + err.Error())
		}
	}(client)

	geminiHistoryRepository := repository.NewGeminiHistoryRepository(db)
	geminiService := service.NewGeminiService(client, geminiHistoryRepository)
	helloWorldService := service.NewHelloWorldService()
	geminiHistoryService := service.NewGeminiHistoryService(geminiHistoryRepository)

	helloWorldController := controller.NewHelloWorldController(helloWorldService)
	geminiController := controller.NewGeminiController(geminiService)
	geminiHistoryController := controller.NewGeminiHistoryController(geminiHistoryService)

	route := router.New(helloWorldController, geminiController, geminiHistoryController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	log.Println("Swagger served on http://localhost:8080/swagger/index.html")

	serverError := server.ListenAndServe()

	log.Panicln(serverError)
}
