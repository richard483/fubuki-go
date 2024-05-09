package repository

import (
	"fubuki-go/model"
)

type GeminiHistoryRepositoryInterface interface {
	Create(history *model.History) error
	GetAll() []model.History
	CreateMany(histories *[]model.History) error
	Update(history *model.History) error
	Delete(id string) error
}
