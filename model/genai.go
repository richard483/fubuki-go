package model

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID           uint `gorm:"primaryKey"`
	UserQuestion string
	ModelAnswer  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ModelSource  string         `gorm:"index, default:'gemini'"`
}
