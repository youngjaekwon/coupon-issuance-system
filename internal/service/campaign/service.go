package campaign

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/campaign"
	"github.com/google/uuid"
)

type Service interface {
	CreateCampaign(ctx context.Context, input *CreateCampaignInput) (*CampaignOutput, error)
	FindCampaign(ctx context.Context, id uuid.UUID) (*CampaignOutput, error)
	ListCampaigns(ctx context.Context, page, limit int) ([]*CampaignOutput, error)
}

type service struct {
	repository repo.Repository
}

func New(repository repo.Repository) Service {
	return &service{repository: repository}
}
