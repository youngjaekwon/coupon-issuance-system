package campaign

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/campaign"
)

type Service interface {
	CreateCampaign(ctx context.Context, input *CreateCampaignInput) (*CreateCampaignOutput, error)
}

type service struct {
	repository repo.Repository
}

func New(repository repo.Repository) Service {
	return &service{repository: repository}
}
