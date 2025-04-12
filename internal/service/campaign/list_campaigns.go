package campaign

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
)

func (s *service) ListCampaigns(ctx context.Context, page, limit int) ([]*CampaignOutput, error) {
	if page < 0 {
		return nil, apperrors.ErrInvalidPage
	}
	if limit < 0 {
		return nil, apperrors.ErrInvalidPageSize
	}
	campaigns, err := s.repository.List(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	var output []*CampaignOutput
	for _, campaign := range campaigns {
		output = append(output, &CampaignOutput{
			ID:         campaign.ID.String(),
			Name:       campaign.Name,
			TotalCount: campaign.TotalCount,
			StartAt:    campaign.StartAt,
			EndAt:      campaign.EndAt,
		})
	}

	return output, nil
}
