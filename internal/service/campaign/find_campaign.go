package campaign

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *service) FindCampaign(ctx context.Context, id uuid.UUID) (*CampaignOutput, error) {
	campaign, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrCampaignNotFound
		}
		return nil, err
	}
	if campaign == nil {
		return nil, apperrors.ErrCampaignNotFound
	}

	couponOutputs := make([]*CouponOutput, len(campaign.Coupons))
	for i, coupon := range campaign.Coupons {
		couponOutputs[i] = &CouponOutput{
			Code:     coupon.Code,
			UserID:   coupon.UserID,
			IssuedAt: coupon.IssuedAt,
		}
	}

	stock := campaign.TotalCount - len(campaign.Coupons)

	return &CampaignOutput{
		ID:         campaign.ID.String(),
		Name:       campaign.Name,
		TotalCount: campaign.TotalCount,
		Stock:      stock,
		StartAt:    campaign.StartAt,
		EndAt:      campaign.EndAt,
		CreatedAt:  campaign.CreatedAt,
		UpdatedAt:  campaign.UpdatedAt,
		Coupons:    couponOutputs,
	}, nil
}
