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
	JWTSecret  []byte
	BcryptCost int
}

type ContextKey string

const ContextKeyDB ContextKey = "db"
const ContextKeyLogger ContextKey = "logger"
const ContextKeyJWTSecret ContextKey = "jwtSecret"
const ContextKeyBcryptCost ContextKey = "bcryptCost"

func Context(config *ContextConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, ContextKeyDB, config.DB)
			ctx = context.WithValue(ctx, ContextKeyLogger, config.Logger)
			ctx = context.WithValue(ctx, ContextKeyJWTSecret, config.JWTSecret)
			ctx = context.WithValue(ctx, ContextKeyBcryptCost, config.BcryptCost)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetDB(r *http.Request) *gorm.DB {
	return r.Context().Value(ContextKeyDB).(*gorm.DB)
}

func GetLogger(r *http.Request) zerolog.Logger {
	return r.Context().Value(ContextKeyLogger).(zerolog.Logger)
}

func GetJWTSecret(r *http.Request) []byte {
	return r.Context().Value(ContextKeyJWTSecret).([]byte)
}

func GetBcryptCost(r *http.Request) int {
	return r.Context().Value(ContextKeyBcryptCost).(int)
}
