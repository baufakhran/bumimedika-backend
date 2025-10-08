package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient(addr, pass string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("[Redis] Connected:", addr)
	return rdb
}

// Set token cache with TTL (in seconds)
func CacheToken(rdb *redis.Client, token string, ttl time.Duration) error {
	return rdb.Set(ctx, token, "valid", ttl).Err()
}

// Check if token exists in cache
func IsTokenCached(rdb *redis.Client, token string) bool {
	val, err := rdb.Get(ctx, token).Result()
	if err == redis.Nil || val != "valid" {
		return false
	}
	return true
}
