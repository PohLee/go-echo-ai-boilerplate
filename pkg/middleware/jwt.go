package middleware

import (
	"strings"

	"github.com/PohLee/go-echo-ai-boilerplate/internal/domain"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/auth"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	jwtService auth.JWTService
}

func NewMiddleware(jwtService auth.JWTService) *Middleware {
	return &Middleware{jwtService: jwtService}
}

func (m *Middleware) JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(401, "missing authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.NewHTTPError(401, "invalid authorization format")
		}

		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(401, "invalid or expired token")
		}

		c.Set(string(domain.ContextKeyUser), claims)
		return next(c)
	}
}
