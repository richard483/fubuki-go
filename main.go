package main

import (
	"fubuki-go/controller"
	"fubuki-go/router"
	"fubuki-go/service"
	"log"

	"net/http"
)

func main() {

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
