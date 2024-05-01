package main

import (
	"fubuki-go/config"
	"fubuki-go/controller"
	"fubuki-go/model"
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

	helloWorldService := service.NewHelloWorldService()
	geminiService := service.NewGeminiService()

	helloWorldController := controller.NewHelloWorldController(helloWorldService)
	geminiController := controller.NewGeminiController(geminiService)

	route := router.New(helloWorldController, geminiController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	serverError := server.ListenAndServe()

	log.Panicln(serverError)
}
