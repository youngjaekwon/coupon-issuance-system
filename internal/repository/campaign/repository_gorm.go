package campaign

import (
	"context"
	"couponIssuanceSystem/internal/models"
	"fmt"
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

func (r *repository) List(ctx context.Context, page, limit int) ([]*models.Campaign, error) {
	if page < 0 || limit < 0 {
		return nil, fmt.Errorf("invalid pagination: page=%d, limit=%d", page, limit)
	}

	var campaigns []*models.Campaign
	var err error
	if limit == 0 {
		err = r.db.WithContext(ctx).Find(&campaigns).Error
	} else {
		err = r.db.WithContext(ctx).Offset(page * limit).Limit(limit).Find(&campaigns).Error
	}
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}
