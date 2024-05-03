package main

import (
	"fubuki-go/config"
	"fubuki-go/controller"
	"fubuki-go/model"
	"fubuki-go/repository"
	"fubuki-go/router"
	"fubuki-go/service"
	"github.com/joho/godotenv"
	"log"

	"net/http"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db := config.NewDbConnection()

	if err := db.AutoMigrate(&model.History{}); err != nil {
		log.Fatalln(err)
	}

	geminiHistoryRepository := repository.NewGeminiHistoryRepository(db)

	helloWorldService := service.NewHelloWorldService()
	geminiService := service.NewGeminiService()
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
