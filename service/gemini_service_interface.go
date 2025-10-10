package service

import (
	"context"
	"fubuki-go/dto/request"
)

type GeminiServiceInterface interface {
	PromptText(prompt *request.PromptText, ctx context.Context) (string, error)
	Chat(prompt *request.PromptText, ctx context.Context) (string, error)
	ResetSession(ctx context.Context) (string, error)
}
