package logger

import (
	"context"
	"net/http"
)

// Logger : the logger struct
type Logger struct{}

// the key type is unexported to avoid context collision
type key int

const contextKey key = iota

// Key : returns the contextKey, only for testing purposes
func Key() key {
	return contextKey
}

// Log : logs an event to stdout
func (l *Logger) Log(e *Event) {
	e.Print()
}

// Middleware : stores the logger object in http context
func Middleware(logger *Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, contextKey, logger)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Get : returns the logger object from http context
func Get(r *http.Request) *Logger {
	return r.Context().Value(contextKey).(*Logger)
}

// New : creates a new logger object
func New() *Logger {
	logger := &Logger{}
	return logger
}
