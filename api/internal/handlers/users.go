package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ravern/gossip/v2/internal/database"
	"github.com/ravern/gossip/v2/internal/jwt"
	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/ravern/gossip/v2/internal/request"
	"github.com/ravern/gossip/v2/internal/response"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Handle    string  `json:"handle" validate:"required"`
	Email     string  `json:"email" validate:"omitempty,email"`
	Password  string  `json:"password" validate:"required"`
	AvatarURL *string `json:"avatar_url"`
}

type registerPayload struct {
	User  database.User `json:"user"`
	Token string        `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	bcryptCost := middleware.GetBcryptCost(r)
	jwtSecret := middleware.GetJWTSecret(r)

	var b RegisterBody
	if err := request.DecodeBody(r, &b); err != nil {
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

	token, err := jwt.Sign(jwtSecret, user)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sign token")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, &registerPayload{
		Token: token,
		User:  user,
	})
}

type LoginBody struct {
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

	var b LoginBody
	if err := request.DecodeBody(r, &b); err != nil {
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

	token, err := jwt.Sign(jwtSecret, user)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sign token")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, &loginPayload{
		Token: token,
		User:  user,
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

	response.WriteJSON(w, &getUserPayload{User: user})
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	response.WriteJSON(w, user)
}

type UpdateCurrentUserBody struct {
	Handle    *string `json:"handle"`
	Email     *string `json:"email" validate:"omitempty,email"`
	Password  *string `json:"password"`
	AvatarURL *string `json:"avatar_url"`
}

func UpdateCurrentUser(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	bcryptCost := middleware.GetBcryptCost(r)
	user := middleware.GetUser(r)

	var b UpdateCurrentUserBody
	if err := request.DecodeBody(r, &b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if b.Handle != nil {
		user.Handle = *b.Handle
	}
	if b.Email != nil {
		user.Email = *b.Email
	}
	if b.Password != nil {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(*b.Password), bcryptCost)
		if err != nil {
			logger.Error().Err(err).Msg("failed to generate password hash")
			response.InternalServerError(w)
			return
		}
		passwordHashString := string(passwordHash)
		user.PasswordHash = &passwordHashString
	}
	if b.AvatarURL != nil {
		user.AvatarURL = b.AvatarURL
	}

	db.Updates(user)

	response.WriteJSON(w, user)
}
