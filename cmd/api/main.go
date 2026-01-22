package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PohLee/go-echo-ai-boilerplate/internal/api_handler/health"
	"github.com/PohLee/go-echo-ai-boilerplate/internal/api_handler/users"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/auth"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/cache"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/config"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/database"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/logger"
	customMiddleware "github.com/PohLee/go-echo-ai-boilerplate/pkg/middleware"
	pkgValidator "github.com/PohLee/go-echo-ai-boilerplate/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

// @title Go Echo AI Boilerplate API
// @version 1.0
// @description This is a production-ready Go Echo boilerplate with AI-First architecture.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// 2. Initialize Logger
	log, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer log.Sync()

	log.Info("Starting Go Echo AI Boilerplate", zap.String("env", cfg.AppEnv))

	// 3. Connect to Database
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Error("Failed to connect to database - proceeding in degraded mode", zap.Error(err))
	}
	_ = db // Reserved for later use

	// 4. Initialize Core Services
	// TODO: Load secret from config
	jwtService := auth.NewJWTService("secret-key-from-config")
	redisClient, err := cache.NewRedisClient(cfg)
	if err != nil {
		log.Error("Failed to connect to Redis", zap.Error(err))
		// Proceed without Redis, or Fatal depending on requirement
	}
	_ = redisClient // Reserved for later use (or passed to cache middleware)

	// 5. Initialize Echo
	e := echo.New()
	e.Validator = pkgValidator.NewValidator()

	// Global Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// New Middlewares
	e.Use(customMiddleware.RequestIDMiddleware())
	e.Use(customMiddleware.SecureMiddleware())
	// Rate Limiter
	e.Use(customMiddleware.RateLimitMiddleware())
	// Performance Logger
	e.Use(customMiddleware.PerformanceLogger(log.Logger))

	// Demo of API Key Middleware (Applied conditionally or to specific group)
	// apiKeyMw := customMiddleware.NewAPIKeyMiddleware(cfg)
	// e.Group("/system", apiKeyMw.Validate)

	// Demo of Cache Middleware
	// cacheMw := customMiddleware.NewCacheMiddleware(redisClient, 5*time.Minute)

	// 6. Register Handlers
	users.UserHandler(e, db, jwtService)
	health.HealthHandler(e, db, redisClient)

	// 7. Health Check
	e.GET("/health/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
			"time":    time.Now().UTC().Format(time.RFC3339),
		})
	})

	// 8. Serve Landing Page
	e.Static("/", "internal/assets") // Serves index.html at / and style.css

	// 9. Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 8. Start Server (Graceful Shutdown)
	go func() {
		if err := e.Start(":" + cfg.AppPort); err != nil && err != http.ErrServerClosed {
			log.Fatal("Shutting down the server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal using channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}
	log.Info("Server exited properly")
}
