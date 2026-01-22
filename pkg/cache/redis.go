package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/PohLee/go-echo-ai-boilerplate/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(cfg *config.Config) (*RedisClient, error) {
	db, err := strconv.Atoi(cfg.RedisDB)
	if err != nil {
		db = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       db,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{Client: rdb}, nil
}
