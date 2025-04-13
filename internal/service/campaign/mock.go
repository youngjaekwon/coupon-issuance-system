package campaign

import (
	"context"
	"couponIssuanceSystem/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCampaignService struct {
	mock.Mock
}

func (m *MockCampaignService) FindCampaign(ctx context.Context, id uuid.UUID) (*CampaignOutput, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	campaign := args.Get(0).(*models.Campaign)
	return &CampaignOutput{
		ID:         campaign.ID.String(),
		Name:       campaign.Name,
		TotalCount: campaign.TotalCount,
		StartAt:    campaign.StartAt,
		EndAt:      campaign.EndAt,
	}, args.Error(1)
}

func (m *MockCampaignService) CreateCampaign(ctx context.Context, input *CreateCampaignInput) (*CampaignOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*CampaignOutput), args.Error(1)
}
func (m *MockCampaignService) ListCampaigns(ctx context.Context, page int, limit int) ([]*CampaignOutput, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]*CampaignOutput), args.Error(1)
}
