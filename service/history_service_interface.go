package service

import (
	"fubuki-go/dto/request"
	"fubuki-go/model"
)

type HistoryServiceInterface interface {
	CreateHistoryData(historyData *request.History) error
	CreateManyHistoryData(historiesData *[]request.History) error
	GetAllHistoryDataByModelSource(modelSource string) *[]model.History
	UpdateHistoryData(historyData *request.UpdateHistory) error
	DeleteHistoryData(id string) error
}
