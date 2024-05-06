package repository

import (
	"fubuki-go/model"
)

type GeminiHistoryRepositoryInterface interface {
	Create(history *model.History) error
	GetAll() []model.History
	Update(history *model.History) error
	Delete(id string) error
}
