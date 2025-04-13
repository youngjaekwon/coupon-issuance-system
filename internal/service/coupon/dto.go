package coupon

import (
	"time"
)

type CouponOutput struct {
	Code       string    `json:"code"`
	CampaignID string    `json:"campaign_id"`
	UserID     string    `json:"user_id"`
	IssuedAt   time.Time `json:"issued_at"`
}
