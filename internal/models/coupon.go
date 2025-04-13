package models

import (
	"github.com/google/uuid"
	"time"
)

type Coupon struct {
	Code       string    `gorm:"primaryKey;size:10"`
	CampaignID uuid.UUID `gorm:"type:uuid;not null;index;index:idx_campaign_user,unique"`
	Campaign   Campaign  `gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID     string    `gorm:"not null;index;index:idx_campaign_user,unique"`
	IssuedAt   time.Time `gorm:"autoCreateTime"`
}
