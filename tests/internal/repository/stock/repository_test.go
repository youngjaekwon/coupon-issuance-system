package stock

import (
	"context"
	"couponIssuanceSystem/internal/apperrors"
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
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "decr-test"
	err := repository.PreWarmStock(ctx, campaignID, 5)
	assert.NoError(t, err)

	count, err := repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, 4, count)

	count, err = repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, 3, count)
}

func TestDecrementStock_NegativeCount(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "decr-test-negative"
	err := repository.PreWarmStock(ctx, campaignID, 0)
	assert.NoError(t, err)

	count, err := repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, -1, count)
}

func TestDecrementStock_StockNotPreWarmed(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "decr-test-not-prewarmed"

	count, err := repository.DecrementStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, -1, count)
}

func TestIncrementStock_Success(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "incr-test"
	totalCount := 5
	err := repository.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	err = repository.IncrementStock(ctx, campaignID)
	assert.NoError(t, err)

	key := repo.StockKey(campaignID)
	valStr, err := redisClient.Get(ctx, key).Result()
	assert.NoError(t, err)

	valInt, err := strconv.Atoi(valStr)
	assert.NoError(t, err)

	expected := totalCount + 1
	assert.Equal(t, expected, valInt)
}

func TestIncrementStock_StockNotPreWarmed(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "incr-test-not-prewarmed"

	err := repository.IncrementStock(ctx, campaignID)
	assert.NoError(t, err)

	key := repo.StockKey(campaignID)
	valStr, err := redisClient.Get(ctx, key).Result()
	assert.NoError(t, err)

	valInt, err := strconv.Atoi(valStr)
	assert.NoError(t, err)

	assert.Equal(t, 1, valInt)
}

func TestRetrieveStock_Success(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "retrieve-test"
	totalCount := 10
	err := repository.PreWarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	count, err := repository.RetrieveStock(ctx, campaignID)
	assert.NoError(t, err)
	assert.Equal(t, totalCount, count)
}

func TestRetrieveStock_StockNotPreWarmed(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "retrieve-test-not-prewarmed"

	count, err := repository.RetrieveStock(ctx, campaignID)
	assert.ErrorIs(t, err, apperrors.ErrStockNotPreWarmed)
	assert.Equal(t, 0, count)
}
