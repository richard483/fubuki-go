package service

import (
	"context"
	"fubuki-go/dto/request"
	"fubuki-go/dto/response_ext"
)

type OllamaServiceInterface interface {
	PromptOllamaText(prompt *request.PromptText, ctx context.Context) (*response_ext.OllamaGenerateResponse, error)
	ChatOllama(prompt *request.PromptText, ctx context.Context) (*response_ext.OllamaChatResponse, error)
	ResetChat() error
}
