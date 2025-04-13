package stock

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	redisClient *redis.Client
}

func New(redisClient *redis.Client) Repository {
	return &repository{redisClient: redisClient}
}

func (r *repository) PreWarmStock(ctx context.Context, campaignID string, totalCount int) error {
	key := StockKey(campaignID)
	return r.redisClient.Set(ctx, key, totalCount, 0).Err()
}

func (r *repository) IsStockPreWarm(ctx context.Context, campaignID string) (bool, error) {
	key := StockKey(campaignID)
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	return val != "", nil
}

func (r *repository) DecrementStock(ctx context.Context, campaignID string) (int64, error) {
	key := StockKey(campaignID)
	val, err := r.redisClient.Decr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}
