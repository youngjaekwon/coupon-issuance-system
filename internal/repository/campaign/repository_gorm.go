package campaign

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

func (r *repository) Create(ctx context.Context, campaign *models.Campaign) error {
	return r.db.WithContext(ctx).Create(campaign).Error
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*models.Campaign, error) {
	var campaign models.Campaign
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&campaign).Error
	if err != nil {
		return nil, err
	}
	return &campaign, nil
}
