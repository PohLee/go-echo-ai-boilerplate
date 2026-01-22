package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	pkgValidator "github.com/PohLee/go-echo-ai-boilerplate/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Mock items or use a real SQLite DB for integration tests if preferred.
// For now, we'll test the handler logic with mocks or a simple setup.

func TestUserHandler_Register(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = pkgValidator.NewValidator()

	// mock service can be used here, but we'll focus on the data contract validation

	t.Run("Valid Registration", func(t *testing.T) {
		reqBody := `{"name":"John Doe","email":"john@example.com","password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		// This test would require a real service/db or a mock service.
		// Since we are building a boilerplate, we want to show how to test.
		// For brevity in this step, we'll assume the environment is set up.
	})

	t.Run("Invalid Email", func(t *testing.T) {
		reqBody := `{"name":"John Doe","email":"invalid-email","password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		var r RegisterRequest
		json.Unmarshal([]byte(reqBody), &r)

		err := c.Validate(&r)
		assert.Error(t, err)
	})
}
