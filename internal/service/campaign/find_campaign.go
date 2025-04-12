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

	return &CampaignOutput{
		ID:         campaign.ID.String(),
		Name:       campaign.Name,
		TotalCount: campaign.TotalCount,
		StartAt:    campaign.StartAt,
		EndAt:      campaign.EndAt,
	}, nil
}
