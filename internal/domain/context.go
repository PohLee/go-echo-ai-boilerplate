package domain

type ContextKey string

const (
	ContextKeyUser      ContextKey = "user"
	ContextKeyRequestID ContextKey = "request_id"
)
