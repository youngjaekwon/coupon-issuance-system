package models

import (
	"github.com/google/uuid"
	"time"
)

type Campaign struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name       string     `gorm:"not null"`
	TotalCount int        `gorm:"not null"`
	StartAt    time.Time  `gorm:"not null"`
	EndAt      *time.Time // optional
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
	Coupons    []Coupon   `gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
