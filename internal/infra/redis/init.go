package redis

import (
	"context"
	"couponIssuanceSystem/internal/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var (
	ctx = context.Background()
)

func Init() *redis.Client {
	addr := config.AppConfig.RedisAddress

	redisClient := redis.NewClient(&redis.Options{
		Addr:         config.AppConfig.RedisAddress,
		Password:     config.AppConfig.RedisPassword,
		DB:           config.AppConfig.RedisDB,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		DialTimeout:  30 * time.Second,
		PoolSize:     100,
		MinIdleConns: 10,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Panicf("Connected to Redis at %s", addr)
	return redisClient
}
