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

	helloWorldController := controller.NewHelloWorldController(helloWorldService)

	route := router.New(helloWorldController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	serverError := server.ListenAndServe()

	log.Panicln(serverError)
}
