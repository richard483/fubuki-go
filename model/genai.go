package model

import (
	"gorm.io/gorm"
	"time"
)

type History struct {
	ID           uint `gorm:"primaryKey"`
	UserQuestion string
	ModelAnswer  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
