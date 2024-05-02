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

func (r *GeminiHistoryRepository) Create(history *model.History) int64 {
	result := r.DB.Create(&history)

	if err := result.Error; err != nil {
		log.Fatalln(err)
	}

	return result.RowsAffected
}

func (r *GeminiHistoryRepository) GetAll() []model.History {
	var histories []model.History
	result := r.DB.Find(&histories)

	if err := result.Error; err != nil {
		log.Fatalln(err)
	}

	return histories
}

func (r *GeminiHistoryRepository) Update(history *model.History) int64 {
	result := r.DB.Save(&history)

	if err := result.Error; err != nil {
		log.Fatalln(err)
	}

	return result.RowsAffected
}

func (r *GeminiHistoryRepository) Delete(id string) int64 {
	result := r.DB.Delete(&model.History{}, id)

	if err := result.Error; err != nil {
		log.Fatalln(err)
	}

	return result.RowsAffected
}
