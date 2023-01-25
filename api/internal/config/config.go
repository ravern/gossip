package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   []byte
	BcryptCost  int
}

func Load(logger zerolog.Logger) (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Warn().Msg("failed to load environment variables from .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		err := fmt.Errorf("missing DATABASE_URL")
		logger.Error().Msg("missing DATABASE_URL")
		return nil, err
	}

	jwtSecretString := os.Getenv("JWT_SECRET")
	if jwtSecretString == "" {
		err := fmt.Errorf("missing JWT_SECRET")
		logger.Error().Msg("missing JWT_SECRET")
		return nil, err
	}
	jwtSecret := []byte(jwtSecretString)

	bcryptCostString := os.Getenv("BCRYPT_COST")
	var bcryptCost int
	if bcryptCostString == "" {
		bcryptCost = 10
	} else {
		bcryptCostUint64, err := strconv.ParseUint(bcryptCostString, 10, 64)
		if err != nil {
			logger.Error().Err(err).Msg("failed to parse BCRYPT_COST")
			return nil, err
		}
		bcryptCost = int(bcryptCostUint64)
	}

	return &Config{
		Port:        port,
		DatabaseURL: databaseURL,
		JWTSecret:   jwtSecret,
		BcryptCost:  bcryptCost,
	}, nil
}
