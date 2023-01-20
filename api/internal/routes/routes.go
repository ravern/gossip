package routes

import (
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
	router.Post("/auth/login", handlers.Login)
	router.With(middleware.SetUser, middleware.ProtectHandler).Get("/users/{id}", handlers.GetUser)
	router.With(middleware.SetUser, middleware.ProtectHandler).Get("/user", handlers.GetCurrentUser)
	router.With(middleware.SetUser, middleware.ProtectHandler).Put("/user", handlers.UpdateCurrentUser)

	router.With(middleware.SetUser, middleware.ProtectHandler).Post("/posts", handlers.CreatePost)
	router.Get("/posts", handlers.GetAllPosts)
	router.Get("/posts/{id}", handlers.GetPost)
}
