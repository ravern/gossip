package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ravern/gossip/v2/internal/handlers"
	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Config struct {
	DB         *gorm.DB
	Logger     zerolog.Logger
	BcryptCost int
}

func Configure(router chi.Router, config *Config) {
	router.Use(middleware.Context(&middleware.ContextConfig{
		DB:         config.DB,
		Logger:     config.Logger,
		BcryptCost: config.BcryptCost,
	}))

	router.Post("/auth/register", handlers.Register)
	router.Post("/auth/login", func(w http.ResponseWriter, r *http.Request) {})
	router.Get("/user", func(w http.ResponseWriter, r *http.Request) {})
	router.Put("/user", func(w http.ResponseWriter, r *http.Request) {})
}