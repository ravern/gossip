package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ravern/gossip/v2/internal/database"

	validate "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type registerBody struct {
	Handle    string `validate:"required"`
	Email     string `validate:"email"`
	Password  string `validate:"required"`
	AvatarURL string
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := getDB(r)
	logger := getLogger(r)
	bcryptCost := getBcryptCost(r)

	var b registerBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request")
		jsonError(w, err, http.StatusBadRequest)
		return
	}

	if err := validate.New().Struct(&b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		jsonError(w, err, http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcryptCost)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate password hash")
		internalServerError(w)
		return
	}

	user := database.User{
		Handle:       b.Handle,
		Email:        b.Email,
		PasswordHash: string(passwordHash),
		AvatarURL:    b.AvatarURL,
	}

	result := db.Create(&user)
	if result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to create user")
		internalServerError(w)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {}

func GetUser(w http.ResponseWriter, r *http.Request) {}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}
