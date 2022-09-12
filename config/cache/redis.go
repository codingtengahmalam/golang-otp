package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

func InitRedis() Redis {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	return &redisClient{
		rdb: rdb,
	}
}

type redisClient struct {
	rdb *redis.Client
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(tx context.Context, key string) (string, error)
}

func (c *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = c.rdb.Set(ctx, key, jsonData, 10*time.Minute).Err()
	return err
}

func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
