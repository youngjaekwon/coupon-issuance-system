package stock

import "context"

type Repository interface {
	PreWarmStock(ctx context.Context, campaignID string, totalCount int) error
	IsStockPreWarm(ctx context.Context, campaignID string) (bool, error)
	DecrementStock(ctx context.Context, campaignID string) (int64, error)
}
