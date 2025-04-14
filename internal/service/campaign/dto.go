package campaign

import (
	"time"
)

type CreateCampaignInput struct {
	Name       string
	TotalCount int
	StartAt    time.Time
	EndAt      *time.Time
}

func (c *CreateCampaignInput) IsValid() bool {
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

type CouponOutput struct {
	Code     string    `json:"code"`
	UserID   string    `json:"user_id"`
	IssuedAt time.Time `json:"issued_at"`
}

type CampaignOutput struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	TotalCount int             `json:"total_count"`
	Stock      int             `json:"stock"`
	StartAt    time.Time       `json:"start_at"`
	EndAt      *time.Time      `json:"end_at,omitempty"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	Coupons    []*CouponOutput `json:"coupons,omitempty"`
}
