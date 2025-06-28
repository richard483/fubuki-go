package impl

import (
	"fubuki-go/model"
	"log/slog"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	*gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db}
}

func (r *HistoryRepository) Create(history *model.History) error {
	result := r.DB.Create(&history)
	return result.Error
}

func (r *HistoryRepository) CreateMany(histories *[]model.History) error {
	result := r.DB.Create(&histories)
	return result.Error
}

func (r *HistoryRepository) GetAllByModelSource(modelSource string) []model.History {
	var histories []model.History
	result := r.DB.Where(&model.History{DeletedAt: gorm.DeletedAt{}, ModelSource: modelSource}).Find(&histories)

	if err := result.Error; err != nil {
		slog.Error("#GetAllByModelSource - error fetching histories", "error", err.Error())
	}

	return histories
}

func (r *HistoryRepository) Update(history *model.History) error {
	result := r.DB.Save(&history)

	return result.Error
}

func (r *HistoryRepository) Delete(id string) error {
	result := r.DB.Delete(&model.History{}, id)

	return result.Error
}
