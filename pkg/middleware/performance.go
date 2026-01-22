package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func PerformanceLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			duration := time.Since(start)
			status := c.Response().Status
			method := c.Request().Method
			path := c.Request().URL.Path

			logger.Info("Request Performance",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", status),
				zap.Duration("latency", duration),
			)

			return err
		}
	}
}
