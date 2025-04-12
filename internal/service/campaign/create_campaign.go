package campaign

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	"couponIssuanceSystem/internal/models"
)

func (s *service) CreateCampaign(ctx context.Context, input *CreateCampaignInput) (*CreateCampaignOutput, error) {
	if !input.IsValid() {
		return nil, apperrors.ErrInvalidCampaignInput
	}

	campaign := &models.Campaign{
		Name:       input.Name,
		TotalCount: input.TotalCount,
		StartAt:    input.StartAt,
	}
	if input.EndAt != nil {
		campaign.EndAt = input.EndAt
	}

	if err := s.repository.Create(ctx, campaign); err != nil {
		return nil, err
	}

	return &CreateCampaignOutput{
		ID:         campaign.ID.String(),
		Name:       campaign.Name,
		TotalCount: campaign.TotalCount,
		StartAt:    campaign.StartAt,
		EndAt:      campaign.EndAt,
	}, nil
}
