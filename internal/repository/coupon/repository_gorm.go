package coupon

import (
	"context"
	"couponIssuanceSystem/internal/models"
	"github.com/google/uuid"
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

func (r *repository) ExistsByUser(ctx context.Context, campaignID uuid.UUID, userID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Coupon{}).
		Where("campaign_id = ? AND user_id = ?", campaignID, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
