package domain

import "errors"

// Standard Application Errors
var (
	ErrNotFound         = errors.New("resource not found")
	ErrInternal         = errors.New("internal server error")
	ErrInvalidRequest   = errors.New("invalid request")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrConflict         = errors.New("resource conflict")
	ErrValidationFailed = errors.New("validation failed")
)
