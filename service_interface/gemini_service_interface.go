package service_interface

import (
	"fubuki-go/dto/request"
)

type GeminiServiceInterface interface {
	PromptText(prompt *request.PromptText) (error, *[]string)
	Chat(prompt *request.PromptText) (error, *[]string)
	ResetSession() (error, string)
}
