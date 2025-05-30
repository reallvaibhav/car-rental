package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
		return nil
	}

	log.Println("Connected to Redis cache")
	return &RedisCache{client: rdb}
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if r.client == nil {
		return fmt.Errorf("redis client not available")
	}

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, json, expiration).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string, dest interface{}) error {
	if r.client == nil {
		return fmt.Errorf("redis client not available")
	}

	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	if r.client == nil {
		return fmt.Errorf("redis client not available")
	}

	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) Close() error {
	if r.client == nil {
		return nil
	}
	return r.client.Close()
}
