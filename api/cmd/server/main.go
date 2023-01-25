package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/ravern/gossip/v2/internal/config"
	"github.com/ravern/gossip/v2/internal/database"
	"github.com/ravern/gossip/v2/internal/routes"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.NewConsoleWriter())

	config, err := config.Load(logger)
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Error().Err(err).Msg("failed to connect to database")
		return err
	}
	logger.Info().Msg("connected to database")
	database.Configure(db)

	router := chi.NewRouter()
	routes.Configure(router, &routes.Config{
		DB:         db,
		Logger:     logger,
		JWTSecret:  config.JWTSecret,
		BcryptCost: config.BcryptCost,
	})

	logger.Info().Str("port", config.Port).Msg("server listening")
	if err := http.ListenAndServe(config.Port, router); err != nil {
		logger.Error().Err(err).Msg("failed to start server")
		return err
	}

	return nil
}
