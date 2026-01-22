package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestIDMiddleware() echo.MiddlewareFunc {
	return middleware.RequestID()
}
