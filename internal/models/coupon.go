package models

import (
	"github.com/google/uuid"
	"time"
)

type Coupon struct {
	Code       string    `json:"code" gorm:"primaryKey;size:10"`
	CampaignID uuid.UUID `json:"campaign_id" gorm:"type:uuid;not null;index;index:idx_campaign_user,unique"`
	Campaign   Campaign  `json:"campaign" gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID     string    `json:"user_id" gorm:"not null;index;index:idx_campaign_user,unique"`
	IssuedAt   time.Time `json:"issued_at" gorm:"autoCreateTime"`
}
