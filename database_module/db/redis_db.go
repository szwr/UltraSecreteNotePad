package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
    *redis.Client
}

func ReturnRedisClient(addr string, password string, db int) *RedisClient {
    c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

    return &RedisClient{c}
}

func (c *RedisClient) AddKeyValue(key string, val string) error {
    ctx := context.Background()

	err := c.Set(ctx, key, val, 0).Err()
    if err != nil {
        return err
    }

    return nil
}

func (c *RedisClient) GetKeyValue(key string) (string, error) {
	ctx := context.Background()

	val, err := c.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
