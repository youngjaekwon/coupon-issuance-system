package coupon

import (
	"context"
	"couponIssuanceSystem/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCouponRepository struct {
	mock.Mock
}

func (m *MockCouponRepository) Create(ctx context.Context, coupon *models.Coupon) (bool, error) {
	args := m.Called(ctx, coupon)
	return args.Bool(0), args.Error(1)
}

func (m *MockCouponRepository) ExistsByUser(ctx context.Context, campaignID uuid.UUID, userID string) (bool, error) {
	args := m.Called(ctx, campaignID, userID)
	return args.Bool(0), args.Error(1)
}
