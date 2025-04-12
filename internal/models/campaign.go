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

func (c *Campaign) IsValid() bool {
	if c.ID == uuid.Nil {
		return false
	}

	if c.Name == "" {
		return false
	}

	if c.TotalCount <= 0 {
		return false
	}

	if c.StartAt.IsZero() || c.StartAt.Before(time.Now()) {
		return false
	}

	return true
}
