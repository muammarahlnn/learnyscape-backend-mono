package data

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
}

type redisClient struct {
	rdb *redis.Client
}

func NewRedisClient(rdb *redis.Client) RedisClient {
	return &redisClient{
		rdb: rdb,
	}
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *redisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	err := r.rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisClient) Delete(ctx context.Context, key string) error {
	err := r.rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
