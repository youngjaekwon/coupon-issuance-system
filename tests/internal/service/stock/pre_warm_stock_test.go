package stock

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/stock"
	svc "couponIssuanceSystem/internal/service/stock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupPreWarmCouponService() (*svc.Service, context.Context, *repo.MockStockRepository) {
	ctx := context.Background()
	stockRepository := new(repo.MockStockRepository)
	service := svc.New(stockRepository)
	return service, ctx, stockRepository
}

func TestPreWarmStock_Success(t *testing.T) {
	service, ctx, stockRepository := setupPreWarmCouponService()
	campaignID := "campaign123"
	totalCount := 100

	stockRepository.On("IsStockPreWarm", ctx, campaignID).Return(false, nil)
	stockRepository.On("PreWarmStock", ctx, campaignID, totalCount).Return(nil)

	err := service.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	stockRepository.AssertExpectations(t)
}

func TestPreWarmStock_AlreadyPreWarm(t *testing.T) {
	service, ctx, stockRepository := setupPreWarmCouponService()
	campaignID := "campaign123"
	totalCount := 100

	stockRepository.On("IsStockPreWarm", ctx, campaignID).Return(true, nil)

	err := service.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	stockRepository.AssertExpectations(t)
}
