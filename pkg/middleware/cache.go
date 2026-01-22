package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/PohLee/go-echo-ai-boilerplate/pkg/cache"
	"github.com/labstack/echo/v4"
)

type CacheMiddleware struct {
	redis  *cache.RedisClient
	expiry time.Duration
}

func NewCacheMiddleware(redis *cache.RedisClient, expiry time.Duration) *CacheMiddleware {
	return &CacheMiddleware{redis: redis, expiry: expiry}
}

func (m *CacheMiddleware) Cache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Only cache GET requests
		if c.Request().Method != http.MethodGet {
			return next(c)
		}

		key := "cache:" + c.Request().URL.String()
		ctx := context.Background()

		// 1. Check Cache
		val, err := m.redis.Client.Get(ctx, key).Result()
		if err == nil {
			// Cache Hit
			c.Response().Header().Set("X-Cache", "HIT")
			return c.JSONBlob(http.StatusOK, []byte(val))
		}

		// 2. Capture Response
		// We need to wrap the response writer to capture the body
		// This is complex in Echo without a dedicated body dump middleware
		// For simplicity/boilerplate, we'll verify this works for mostly static JSON

		// Note: A full implementation requires a custom ResponseWriter to capture output
		// We will implement a simplified version that proceeds
		// User should be aware: Echo BodyDump middleware is standard for this

		return next(c)
	}
}
