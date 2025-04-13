package stock

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
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

func (r *repository) DecrementStock(ctx context.Context, campaignID string) (int, error) {
	key := StockKey(campaignID)
	val, err := r.redisClient.Decr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

func (r *repository) IncrementStock(ctx context.Context, campaignID string) error {
	key := StockKey(campaignID)
	_, err := r.redisClient.Incr(ctx, key).Result()
	return err
}

func (r *repository) RetrieveStock(ctx context.Context, campaignID string) (int, error) {
	key := StockKey(campaignID)
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, apperrors.ErrStockNotPreWarmed
		}
		return 0, err
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return valInt, nil
}
