package stock

import (
	"context"
	repo "couponIssuanceSystem/internal/repository/stock"
	test_redis "couponIssuanceSystem/tests/infra/redis"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestPrewarmStock_Success(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "test-campaign-id"
	totalCount := 100

	err := repository.PrewarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	key := repo.StockKey(campaignID)
	valStr, err := redisClient.Get(ctx, key).Result()
	assert.NoError(t, err)

	valInt, err := strconv.Atoi(valStr)
	assert.NoError(t, err)

	assert.Equal(t, totalCount, valInt)
}

func TestIsStockPrewarmed_Success(t *testing.T) {
	ctx := context.Background()
	redisClient, mr := test_redis.SetupTestRedisClient(ctx)
	defer mr.Close()
	repository := repo.New(redisClient)

	campaignID := "test-campaign-id"
	totalCount := 100

	err := repository.PrewarmStock(ctx, campaignID, totalCount)
	assert.NoError(t, err)

	isPrewarm, err := repository.IsStockPrewarm(ctx, campaignID)
	assert.NoError(t, err)
	assert.True(t, isPrewarm)
}
