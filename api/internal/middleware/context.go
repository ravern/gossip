package middleware

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type ContextConfig struct {
	DB         *gorm.DB
	Logger     zerolog.Logger
	BcryptCost int
}

type ContextKey string

const ContextKeyDB ContextKey = "db"
const ContextKeyLogger ContextKey = "logger"
const ContextKeyBcryptCost ContextKey = "bcryptCost"

func Context(config *ContextConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, ContextKeyDB, config.DB)
			ctx = context.WithValue(ctx, ContextKeyLogger, config.Logger)
			ctx = context.WithValue(ctx, ContextKeyBcryptCost, config.BcryptCost)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
