package service_interface

import (
	"fubuki-go/dto/request"
	"fubuki-go/model"
)

type GeminiHistoryServiceInterface interface {
	CreateHistoryData(historyData *request.GeminiHistory) error
	CreateManyHistoryData(historiesData *[]request.GeminiHistory) error
	GetAllHistoryData() *[]model.History
	UpdateHistoryData(historyData *request.UpdateGeminiHistory) error
	DeleteHistoryData(id string) error
}
