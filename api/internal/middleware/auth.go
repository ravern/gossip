package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ravern/gossip/v2/internal/database"
	"github.com/ravern/gossip/v2/internal/jwt"
	"github.com/ravern/gossip/v2/internal/response"
)

type AuthContextKey string

const AuthContextKeyUser AuthContextKey = "user"

func SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := GetDB(r)
		jwtSecret := GetJWTSecret(r)

		authorizationString := r.Header.Get("Authorization")
		if authorizationString == "" {
			next.ServeHTTP(w, r)
			return
		}
		if !strings.HasPrefix(authorizationString, "Bearer") {
			response.JSONError(w, errors.New("only Bearer supported"), http.StatusUnauthorized)
			return
		}

		authorizationComponents := strings.Fields(authorizationString)
		if len(authorizationComponents) != 2 {
			response.JSONError(w, errors.New("invalid Authorization header"), http.StatusUnauthorized)
			return
		}

		claims, err := jwt.Parse(jwtSecret, authorizationComponents[1])
		if err != nil {
			response.JSONError(w, err, http.StatusUnauthorized)
			return
		}

		var user *database.User
		if result := db.Where("id = ?", claims.ID).First(&user); result.Error != nil {
			response.JSONError(w, result.Error, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), AuthContextKeyUser, user)))
	})
}

func ProtectHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Context().Value(AuthContextKeyUser).(*database.User)
		if !ok {
			response.JSONError(w, errors.New("login required"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetUser(r *http.Request) *database.User {
	return r.Context().Value(AuthContextKeyUser).(*database.User)
}
