package handlers

import (
	"net/http"

	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func getDB(r *http.Request) *gorm.DB {
	return r.Context().Value(middleware.ContextKeyDB).(*gorm.DB)
}

func getLogger(r *http.Request) zerolog.Logger {
	return r.Context().Value(middleware.ContextKeyLogger).(zerolog.Logger)
}

func getBcryptCost(r *http.Request) int {
	return r.Context().Value(middleware.ContextKeyBcryptCost).(int)
}
