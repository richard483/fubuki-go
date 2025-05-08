package service_interface

import (
	"fubuki-go/dto/request"
)

type OllamaServiceInterface interface {
	Generate(prompt *request.GeminiText) (error, *[]string)
}
