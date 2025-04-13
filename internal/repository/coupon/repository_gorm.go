package coupon

import (
	"context"
	"couponIssuanceSystem/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, coupon *models.Coupon) error {
	return r.db.WithContext(ctx).Create(coupon).Error
}
