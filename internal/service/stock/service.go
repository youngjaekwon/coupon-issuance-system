package stock

import (
	"context"
	campaignrepo "couponIssuanceSystem/internal/repository/campaign"
	repo "couponIssuanceSystem/internal/repository/stock"
	"time"
)

type Service interface {
	PreWarmStock(ctx context.Context, start, end time.Time) error
}

type service struct {
	repository         repo.Repository
	campaignRepository campaignrepo.Repository
}

func New(repository repo.Repository, campaignRepository campaignrepo.Repository) Service {
	return &service{
		repository:         repository,
		campaignRepository: campaignRepository,
	}
}
