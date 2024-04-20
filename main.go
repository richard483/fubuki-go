package main

import (
	"fubuki-go/controller"
	"fubuki-go/router"
	"fubuki-go/service_impl"
	"log"

	"net/http"
)

func main() {

	helloWorldServiceImpl := service_impl.HelloWorldServiceImpl{}

	helloWorldController := controller.HelloWorldController{HelloWorldService: helloWorldServiceImpl}

	route := router.New(helloWorldController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	serverError := server.ListenAndServe()

	log.Panicln(serverError)
}
