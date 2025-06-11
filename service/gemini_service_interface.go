package service

import (
	"fubuki-go/dto/request"
)

type GeminiServiceInterface interface {
	PromptText(prompt *request.PromptText) (*[]string, error)
	Chat(prompt *request.PromptText) (*[]string, error)
	ResetSession() (string, error)
}
