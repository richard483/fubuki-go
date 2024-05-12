package service_interface

import (
	"fubuki-go/dto/request"
	request2 "fubuki-go/dto/response_ext"
)

type GeminiServiceInterface interface {
	PromptText(prompt *request.GeminiText) (error, *[]string)
	Chat(prompt *request.GeminiText) (error, *[]string)
	TuneModel() (error, *request2.GeminiTunedModel)
	ResetSession() (error, string)
}
