package health

import (
	"context"
	"net/http"
	"time"

	"github.com/PohLee/go-echo-ai-boilerplate/pkg/cache"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StatusResponse struct {
	Status    string            `json:"status"`
	Timestamp string            `json:"timestamp"`
	Services  map[string]string `json:"services"`
}

func HealthHandler(e *echo.Echo, db *gorm.DB, redis *cache.RedisClient) {
	e.GET("/api/status", func(c echo.Context) error {
		status := "ok"
		services := make(map[string]string)

		// Check DB
		if db == nil {
			services["database"] = "unreachable (not connected)"
			status = "error"
		} else {
			sqlDB, err := db.DB()
			if err != nil {
				services["database"] = "unreachable"
				status = "error"
			} else {
				if err := sqlDB.Ping(); err != nil {
					services["database"] = "error"
					status = "error"
				} else {
					services["database"] = "connected"
				}
			}
		}

		// Check Redis
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if redis != nil {
			if err := redis.Client.Ping(ctx).Err(); err != nil {
				services["redis"] = "error"
				status = "error"
			} else {
				services["redis"] = "connected"
			}
		} else {
			services["redis"] = "not_configured"
		}

		return c.JSON(http.StatusOK, StatusResponse{
			Status:    status,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Services:  services,
		})
	})
}
