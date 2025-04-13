package stock

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/stock"
	test_redis "couponIssuanceSystem/tests/infra/redis"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestPreWarmStock_Success(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "test-campaign-id"
	totalCount := 100

	err := repository.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	key := repo.StockKey(campaignID)
	valStr, err := redisClient.Get(ctx, key).Result()
	assert.NoError(t, err)

	valInt, err := strconv.Atoi(valStr)
	assert.NoError(t, err)

	assert.Equal(t, totalCount, valInt)
}

func TestIsStockPreWarmed_Success(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "test-campaign-id"
	totalCount := 100

	err := repository.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	isPreWarm, err := repository.IsStockPreWarm(ctx, campaignID)
	assert.NoError(t, err)
	assert.True(t, isPreWarm)
}

func TestDecrementStock_Success(t *testing.T) {
	ctx := context.Background()
	rdb, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(rdb)

	campaignID := "decr-test"
	err := repository.PreWarmStock(ctx, campaignID, 5)
	assert.NoError(t, err)

	count, err := repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, int64(4), count)

	count, err = repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), count)
}

func TestDecrementStock_NegativeCount(t *testing.T) {
	ctx := context.Background()
	rdb, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(rdb)

	campaignID := "decr-test-negative"
	err := repository.PreWarmStock(ctx, campaignID, 0)
	assert.NoError(t, err)

	count, err := repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, int64(-1), count)
}

func TestDecrementStock_StockNotPreWarmed(t *testing.T) {
	ctx := context.Background()
	rdb, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(rdb)

	campaignID := "decr-test-not-prewarmed"

	count, err := repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, int64(-1), count)
}
