package middleware

import (
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/config"
	"github.com/labstack/echo/v4"
)

type APIKeyMiddleware struct {
	validKeys map[string]bool
}

func NewAPIKeyMiddleware(cfg *config.Config) *APIKeyMiddleware {
	// TODO: Load keys from secure source, for now using env/config
	// Assuming config might have a comma-separated list or single key
	keys := make(map[string]bool)
	if cfg.AppEnv == "local" {
		keys["dev-key-123"] = true
	}
	// Add logic to load real keys

	return &APIKeyMiddleware{validKeys: keys}
}

func (m *APIKeyMiddleware) Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Request().Header.Get("X-API-Key")
		if key == "" {
			return echo.NewHTTPError(401, "missing api key")
		}

		// Constant time comparison not strictly needed for map lookup but good practice if comparing secrets directly
		// Here we just check existence
		if !m.validKeys[key] {
			return echo.NewHTTPError(403, "invalid api key")
		}

		return next(c)
	}
}
