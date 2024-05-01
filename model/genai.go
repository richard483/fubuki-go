package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type History struct {
	ID           uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	UserQuestion string
	ModelAnswer  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
