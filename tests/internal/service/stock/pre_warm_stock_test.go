package stock

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/stock"
	svc "couponIssuanceSystem/internal/service/stock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupPreWarmCouponService() (svc.Service, context.Context, *repo.MockStockRepository) {
	ctx := context.Background()
	stockRepository := new(repo.MockStockRepository)
	service := svc.New(stockRepository)
	return service, ctx, stockRepository
}

func TestPreWarmStock_Success(t *testing.T) {
	service, ctx, stockRepository := setupPreWarmCouponService()
	campaignID := uuid.New()
	totalCount := 100

	stockRepository.On("IsStockPreWarm", ctx, campaignID.String()).Return(false, nil)
	stockRepository.On("PreWarmStock", ctx, campaignID.String(), totalCount).Return(nil)

	err := service.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	stockRepository.AssertExpectations(t)
}

func TestPreWarmStock_AlreadyPreWarm(t *testing.T) {
	service, ctx, stockRepository := setupPreWarmCouponService()
	campaignID := uuid.New()
	totalCount := 100

	stockRepository.On("IsStockPreWarm", ctx, campaignID.String()).Return(true, nil)

	err := service.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	stockRepository.AssertExpectations(t)
}
