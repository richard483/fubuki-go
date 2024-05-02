package repository

import (
	"fubuki-go/model"
)

type GeminiHistoryRepositoryInterface interface {
	Create(history *model.History) int64
	GetAll() []model.History
	Update(history *model.History) int64
	Delete(id string) int64
}
