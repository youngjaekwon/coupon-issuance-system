package coupon

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	"couponIssuanceSystem/internal/models"
	repo "couponIssuanceSystem/internal/repository/coupon"
	stockrepo "couponIssuanceSystem/internal/repository/stock"
	campaignsvc "couponIssuanceSystem/internal/service/campaign"
	svc "couponIssuanceSystem/internal/service/coupon"
	"couponIssuanceSystem/internal/utils/couponcode"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

func setupIssueCouponService() (svc.Service, context.Context, *repo.MockCouponRepository, *campaignsvc.MockCampaignService, *stockrepo.MockStockRepository, *couponcode.MockCodeGenerator) {
	ctx := context.Background()

	couponRepository := new(repo.MockCouponRepository)
	campaignService := new(campaignsvc.MockCampaignService)
	stockRepository := new(stockrepo.MockStockRepository)
	codeGenerator := new(couponcode.MockCodeGenerator)

	service := svc.New(couponRepository, campaignService, stockRepository, codeGenerator)
	return service, ctx, couponRepository, campaignService, stockRepository, codeGenerator
}

func TestIssueCoupon_Success(t *testing.T) {
	service, ctx, couponRepository, campaignService, stockRepository, codeGenerator := setupIssueCouponService()

	campaignID := uuid.New()
	userID := "user123"
	campaign := &models.Campaign{
		ID:         campaignID,
		Name:       "Test Campaign",
		StartAt:    time.Now().Add(-time.Hour),
		TotalCount: 100,
	}

	campaignService.On("FindCampaign", mock.Anything, campaignID).Return(campaign, nil)
	stockRepository.On("DecrementStock", mock.Anything, campaignID.String()).Return(99, nil)
	codeGenerator.On("Generate").Return("COUPON_CODE").Once()
	couponRepository.On("Create", mock.Anything, mock.AnythingOfType("*models.Coupon")).Return(true, nil)

	code, err := service.IssueCoupon(ctx, campaignID, userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, code)
	assert.Equal(t, "COUPON_CODE", code.Code)

	campaignService.AssertExpectations(t)
	couponRepository.AssertExpectations(t)
	stockRepository.AssertExpectations(t)
	codeGenerator.AssertExpectations(t)
}

func TestIssueCoupon_CampaignNotFound(t *testing.T) {
	service, ctx, couponRepository, campaignService, stockRepository, codeGenerator := setupIssueCouponService()

	campaignID := uuid.New()
	userID := "user123"

	campaignService.On("FindCampaign", mock.Anything, campaignID).Return(nil, apperrors.ErrCampaignNotFound)

	code, err := service.IssueCoupon(ctx, campaignID, userID)
	assert.Error(t, err)
	assert.Empty(t, code)

	campaignService.AssertExpectations(t)
	couponRepository.AssertExpectations(t)
	stockRepository.AssertExpectations(t)
	couponRepository.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	codeGenerator.AssertNotCalled(t, "Generate")
}

func TestIssueCoupon_CampaignNotStarted(t *testing.T) {
	service, ctx, couponRepository, campaignService, stockRepository, codeGenerator := setupIssueCouponService()

	campaignID := uuid.New()
	userID := "user123"
	campaign := &models.Campaign{
		ID:         campaignID,
		Name:       "Test Campaign",
		StartAt:    time.Now().Add(time.Hour),
		TotalCount: 100,
	}

	campaignService.On("FindCampaign", mock.Anything, campaignID).Return(campaign, nil)

	code, err := service.IssueCoupon(ctx, campaignID, userID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrCampaignNotStarted))
	assert.Empty(t, code)

	campaignService.AssertExpectations(t)
	couponRepository.AssertExpectations(t)
	stockRepository.AssertExpectations(t)
	couponRepository.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	codeGenerator.AssertNotCalled(t, "Generate")
}

func TestIssueCoupon_CampaignEnded(t *testing.T) {
	service, ctx, couponRepository, campaignService, stockRepository, codeGenerator := setupIssueCouponService()

	campaignID := uuid.New()
	userID := "user123"
	endAt := time.Now().Add(-time.Minute)
	campaign := &models.Campaign{
		ID:         campaignID,
		Name:       "Test Campaign",
		StartAt:    time.Now().Add(-time.Hour),
		EndAt:      &endAt,
		TotalCount: 100,
	}

	campaignService.On("FindCampaign", mock.Anything, campaignID).Return(campaign, nil)

	code, err := service.IssueCoupon(ctx, campaignID, userID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrCampaignEnded))
	assert.Empty(t, code)

	campaignService.AssertExpectations(t)
	couponRepository.AssertExpectations(t)
	stockRepository.AssertExpectations(t)
	couponRepository.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	codeGenerator.AssertNotCalled(t, "Generate")
}

func TestIssueCoupon_SoldOut(t *testing.T) {
	service, ctx, couponRepository, campaignService, stockRepository, codeGenerator := setupIssueCouponService()

	campaignID := uuid.New()
	userID := "user123"
	campaign := &models.Campaign{
		ID:         campaignID,
		Name:       "Test Campaign",
		StartAt:    time.Now().Add(-time.Hour),
		TotalCount: 100,
	}

	campaignService.On("FindCampaign", mock.Anything, campaignID).Return(campaign, nil)
	stockRepository.On("DecrementStock", mock.Anything, campaignID.String()).Return(-1, nil)
	stockRepository.On("IncrementStock", mock.Anything, campaignID.String()).Return(nil)

	code, err := service.IssueCoupon(ctx, campaignID, userID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrCampaignSoldOut))
	assert.Empty(t, code)

	campaignService.AssertExpectations(t)
	couponRepository.AssertExpectations(t)
	stockRepository.AssertExpectations(t)
	couponRepository.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	codeGenerator.AssertNotCalled(t, "Generate")
}

func TestIssueCoupon_DuplicatedCode(t *testing.T) {
	service, ctx, couponRepository, campaignService, stockRepository, codeGenerator := setupIssueCouponService()

	campaignID := uuid.New()
	userID := "user123"
	campaign := &models.Campaign{
		ID:         campaignID,
		Name:       "Test Campaign",
		StartAt:    time.Now().Add(-time.Hour),
		TotalCount: 100,
	}

	campaignService.On("FindCampaign", mock.Anything, campaignID).Return(campaign, nil)
	stockRepository.On("DecrementStock", mock.Anything, campaignID.String()).Return(99, nil)

	codeGenerator.On("Generate").Return("DUPLICATED_CODE").Once()
	codeGenerator.On("Generate").Return("NEW_CODE").Once()

	couponRepository.
		On("Create", mock.Anything, mock.MatchedBy(func(c *models.Coupon) bool {
			return c.Code == "DUPLICATED_CODE"
		})).Return(false, gorm.ErrDuplicatedKey)

	couponRepository.
		On("Create", mock.Anything, mock.MatchedBy(func(c *models.Coupon) bool {
			return c.Code == "NEW_CODE"
		})).Return(true, nil)

	code, err := service.IssueCoupon(ctx, campaignID, userID)
	assert.NoError(t, err)
	assert.Equal(t, "NEW_CODE", code.Code)

	campaignService.AssertExpectations(t)
	couponRepository.AssertExpectations(t)
	stockRepository.AssertExpectations(t)
	codeGenerator.AssertExpectations(t)
}
