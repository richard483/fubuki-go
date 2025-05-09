package service_interface

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response_ext"
)

type OllamaServiceInterface interface {
	Generate(prompt *request.PromptText) (*response_ext.OllamaGenerateResponse, error)
}
