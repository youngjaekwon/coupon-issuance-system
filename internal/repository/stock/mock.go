package stock

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockStockRepository struct {
	mock.Mock
}

func (m *MockStockRepository) PreWarmStock(ctx context.Context, campaignID string, totalCount int) error {
	args := m.Called(ctx, campaignID, totalCount)
	return args.Error(0)
}

func (m *MockStockRepository) IsStockPreWarm(ctx context.Context, campaignID string) (bool, error) {
	args := m.Called(ctx, campaignID)
	return args.Bool(0), args.Error(1)
}

func (m *MockStockRepository) DecrementStock(ctx context.Context, campaignID string) (int, error) {
	args := m.Called(ctx, campaignID)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockStockRepository) IncrementStock(ctx context.Context, campaignID string) error {
	args := m.Called(ctx, campaignID)
	return args.Error(0)
}

func (m *MockStockRepository) RetrieveStock(ctx context.Context, campaignID string) (int, error) {
	args := m.Called(ctx, campaignID)
	return args.Int(0), args.Error(1)
}
