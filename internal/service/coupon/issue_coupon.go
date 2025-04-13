package coupon

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	"couponIssuanceSystem/internal/models"
	"github.com/google/uuid"
	"time"
)

func (s *service) IssueCoupon(ctx context.Context, campaignID uuid.UUID, userID string) (CouponOutput, error) {
	campaign, err := s.campaignService.FindCampaign(ctx, campaignID)
	if err != nil {
		return CouponOutput{}, err
	}
	if campaign == nil {
		return CouponOutput{}, apperrors.ErrCampaignNotFound
	}

	now := time.Now().UTC()
	if now.Before(campaign.StartAt.UTC()) {
		return CouponOutput{}, apperrors.ErrCampaignNotStarted
	}
	if campaign.EndAt != nil && now.After((*campaign.EndAt).UTC()) {
		return CouponOutput{}, apperrors.ErrCampaignEnded
	}

	stock, err := s.stockRepository.DecrementStock(ctx, campaign.ID)
	if err != nil {
		return CouponOutput{}, err
	}
	if stock < 0 {
		_ = s.stockRepository.IncrementStock(ctx, campaign.ID)
		return CouponOutput{}, apperrors.ErrCampaignSoldOut
	}

	maxAttempts := 3
	for i := 0; i < maxAttempts; i++ {
		code := s.codeGenerator.Generate()
		coupon := &models.Coupon{
			Code:       code,
			CampaignID: campaignID,
			UserID:     userID,
		}
		created, err := s.repository.Create(ctx, coupon)

		if !created && err == nil {
			_ = s.stockRepository.IncrementStock(ctx, campaign.ID)
			return CouponOutput{}, apperrors.ErrUserAlreadyIssued
		}
		if created {
			return CouponOutput{
				Code:       coupon.Code,
				CampaignID: coupon.CampaignID.String(),
				UserID:     coupon.UserID,
				IssuedAt:   coupon.IssuedAt,
			}, nil
		}
	}
	_ = s.stockRepository.IncrementStock(ctx, campaign.ID)
	return CouponOutput{}, apperrors.ErrCouponCodeConflict
}
