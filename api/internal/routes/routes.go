package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ravern/gossip/v2/internal/handlers"
	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Config struct {
	DB         *gorm.DB
	Logger     zerolog.Logger
	JWTSecret  []byte
	BcryptCost int
}

func Configure(router chi.Router, config *Config) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.Context(&middleware.ContextConfig{
		DB:         config.DB,
		Logger:     config.Logger,
		JWTSecret:  config.JWTSecret,
		BcryptCost: config.BcryptCost,
	}))

	router.Post("/auth/register", handlers.Register)
	router.Post("/auth/login", handlers.Login)

	router.With(middleware.SetUser, middleware.ProtectHandler).Get("/users/{id}", handlers.GetUser)

	router.With(middleware.SetUser, middleware.ProtectHandler).Get("/user", handlers.GetCurrentUser)
	router.With(middleware.SetUser, middleware.ProtectHandler).Put("/user", handlers.UpdateCurrentUser)

	router.Route("/posts", func(router chi.Router) {
		router.With(middleware.SetUser, middleware.ProtectHandler).Post("/", handlers.CreatePost)
		router.With(middleware.SetUser, middleware.ProtectHandler).Put("/{id}", handlers.UpdatePost)
		router.With(middleware.SetUser, middleware.ProtectHandler).Delete("/{id}", handlers.DeletePost)
		router.Get("/", handlers.GetAllPosts)
		router.Get("/{id}", handlers.GetPost)
		router.With(middleware.SetUser, middleware.ProtectHandler).Post("/{id}/likes", handlers.LikePost)
	})

	router.Route("/posts/{postID}/comments", func(router chi.Router) {
		router.With(middleware.SetUser, middleware.ProtectHandler).Post("/", handlers.CreateComment)
		router.With(middleware.SetUser, middleware.ProtectHandler).Put("/{id}", handlers.UpdateComment)
		router.With(middleware.SetUser, middleware.ProtectHandler).Delete("/{id}", handlers.DeleteComment)
		router.Get("/", handlers.GetAllComments)
		router.Get("/{id}", handlers.GetComment)
		router.With(middleware.SetUser, middleware.ProtectHandler).Post("/{id}/likes", handlers.LikeComment)
	})
}
