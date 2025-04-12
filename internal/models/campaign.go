package models

import (
	"github.com/google/uuid"
	"time"
)

type Campaign struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	TotalCount int       `json:"total_count" gorm:"not null"`
	StartAt    time.Time `json:"start_at" gorm:"not null"`
	EndAt      time.Time `json:"end_at"` // optional
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
