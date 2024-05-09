package repository

import (
	"fubuki-go/model"
	"gorm.io/gorm"
	"log"
)

type GeminiHistoryRepository struct {
	*gorm.DB
}

func NewGeminiHistoryRepository(db *gorm.DB) *GeminiHistoryRepository {
	return &GeminiHistoryRepository{db}
}

func (r *GeminiHistoryRepository) Create(history *model.History) error {
	result := r.DB.Create(&history)
	return result.Error
}

func (r *GeminiHistoryRepository) CreateMany(histories *[]model.History) error {
	result := r.DB.Create(&histories)
	return result.Error
}

func (r *GeminiHistoryRepository) GetAll() []model.History {
	var histories []model.History
	result := r.DB.Find(&histories)

	if err := result.Error; err != nil {
		log.Fatalln(err)
	}

	return histories
}

func (r *GeminiHistoryRepository) Update(history *model.History) error {
	result := r.DB.Save(&history)

	return result.Error
}

func (r *GeminiHistoryRepository) Delete(id string) error {
	result := r.DB.Delete(&model.History{}, id)

	return result.Error
}
