package service_interface

import (
	"fubuki-go/dto/request"
)

type GeminiServiceInterface interface {
	PromptText(prompt *request.GeminiText) (error, *[]string)
	Chat(prompt *request.GeminiText) (error, *[]string)
	ResetSession() (error, string)
}
