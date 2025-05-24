package service

import (
	"fubuki-go/dto/request"
	"fubuki-go/dto/response_ext"
)

type OllamaServiceInterface interface {
	PromptOllamaText(prompt *request.PromptText) (*response_ext.OllamaGenerateResponse, error)
	ChatOllama(prompt *request.PromptText) (*response_ext.OllamaChatResponse, error)
	ResetChat() error
}
