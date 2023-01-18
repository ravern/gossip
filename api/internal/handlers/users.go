package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ravern/gossip/v2/internal/database"
	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/ravern/gossip/v2/internal/response"
	"gorm.io/gorm"

	validate "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type registerBody struct {
	Handle    string  `json:"handle" validate:"required"`
	Email     string  `json:"email" validate:"email"`
	Password  string  `json:"password" validate:"required"`
	AvatarURL *string `json:"avatar_url"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	bcryptCost := middleware.GetBcryptCost(r)

	var b registerBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := validate.New().Struct(&b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcryptCost)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate password hash")
		response.InternalServerError(w)
		return
	}
	passwordHashString := string(passwordHash)

	user := database.User{
		Handle:       b.Handle,
		Email:        b.Email,
		PasswordHash: &passwordHashString,
		AvatarURL:    b.AvatarURL,
	}

	if result := db.Create(&user); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to create user")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, user)
}

type loginBody struct {
	HandleOrEmail string `json:"handle_or_email" validate:"required"`
	Password      string `json:"password" validate:"required"`
}

type loginPayload struct {
	User  database.User `json:"user"`
	Token string        `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	jwtSecret := middleware.GetJWTSecret(r)

	var b loginBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := validate.New().Struct(&b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	var user database.User
	if result := db.Where("handle = ? OR email = ?", b.HandleOrEmail, b.HandleOrEmail).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Return 400 here instead of 404 to prevent basic handle enumeration.
			err := errors.New("invalid credentials")
			logger.Warn().Err(err).Msg("bad request")
			response.JSONError(w, err, http.StatusBadRequest)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch user")
			response.InternalServerError(w)
		}
		return
	}

	if user.PasswordHash == nil {
		err := errors.New("user cannot use password-based login")
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(b.Password)); err != nil {
		err = errors.New("invalid credentials")
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{ID: user.ID.String()})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Error().Err(err).Msg("failed to sign token")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, &loginPayload{
		User:  user,
		Token: tokenString,
	})
}

type getUserPayload struct {
	database.User
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)

	id := chi.URLParam(r, "id")

	var user database.User
	if result := db.Where("id = ?", id).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("user not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch user")
			response.InternalServerError(w)
		}
		return
	}

	response.WriteJSON(w, &getUserPayload{user})
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	response.WriteJSON(w, user)
}

func UpdateCurrentUser(w http.ResponseWriter, r *http.Request) {}
