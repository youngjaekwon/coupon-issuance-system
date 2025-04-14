package stock

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/stock"
	"github.com/google/uuid"
)

type Service interface {
	PreWarmStock(ctx context.Context, campaignID uuid.UUID, totalCount int) error
}

type service struct {
	repository repo.Repository
}

func New(repository repo.Repository) Service {
	return &service{
		repository: repository,
	}
}
