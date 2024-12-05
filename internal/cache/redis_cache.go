package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

// RedisCache struct
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new Redis client instance
func NewRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Change this if you're using a different Redis instance
		Password: "",               // No password set
		DB:       0,                // Default DB
	})

	// Check if the Redis server is available
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil
	}

	return &RedisCache{client: client}
}

// SetCache sets a key-value pair in Redis with an expiration time
func (r *RedisCache) SetCache(key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache: %v", err)
	}
	return nil
}

// GetCache retrieves a value from Redis by key
func (r *RedisCache) GetCache(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // No data found for the given key
	}
	if err != nil {
		return "", fmt.Errorf("failed to get cache: %v", err)
	}
	return val, nil
}
