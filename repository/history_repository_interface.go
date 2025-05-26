package repository

import (
	"fubuki-go/model"
)

type HistoryRepositoryInterface interface {
	Create(history *model.History) error
	GetAllByModelSource(modelSource string) []model.History
	CreateMany(histories *[]model.History) error
	Update(history *model.History) error
	Delete(id string) error
}
