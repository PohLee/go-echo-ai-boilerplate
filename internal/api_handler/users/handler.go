package users

import (
	"net/http"

	"github.com/PohLee/go-echo-ai-boilerplate/pkg/auth"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

// Initializer for the module
func UserHandler(e *echo.Echo, db *gorm.DB, jwtService auth.JWTService) {
	repo := UserRepository(db)
	svc := UserService(repo, jwtService)
	handler := &Handler{service: svc}

	g := e.Group("/users")
	g.POST("", handler.Register)    // POST /users (Register)
	g.POST("/login", handler.Login) // POST /users/login
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body RegisterRequest true "User Registration Info"
// @Success 201 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Router /users [post]
func (h *Handler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	res, err := h.service.Register(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

// Login godoc
// @Summary User Login
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept  json
// @Produce  json
// @Param login body LoginRequest true "User Login Credentials"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} map[string]string
// @Router /users/login [post]
func (h *Handler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	res, err := h.service.Login(&req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
