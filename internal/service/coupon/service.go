package coupon

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/coupon"
	stockrepo "couponIssuanceSystem/internal/repository/stock"
	campaignsvc "couponIssuanceSystem/internal/service/campaign"
	"couponIssuanceSystem/internal/utils/couponcode"
	"github.com/google/uuid"
)

type Service interface {
	IssueCoupon(ctx context.Context, campaignID uuid.UUID, userID string) (CouponOutput, error)
}

type service struct {
	repository      repo.Repository
	campaignService campaignsvc.Service
	stockRepository stockrepo.Repository
	codeGenerator   couponcode.Generator
}

func New(repository repo.Repository, campaignService campaignsvc.Service, stockRepository stockrepo.Repository, codeGenerator couponcode.Generator) Service {
	return &service{
		repository:      repository,
		campaignService: campaignService,
		stockRepository: stockRepository,
		codeGenerator:   codeGenerator,
	}
}
