package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ravern/gossip/v2/internal/database"
)

type Claims struct {
	jwt.RegisteredClaims
	Role database.UserRole `json:"role"`
}

func Sign(jwtSecret []byte, user database.User) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: user.ID.String(),
		},
		Role: database.NormalRole,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func Parse(jwtSecret []byte, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token provided")
	}

	return claims, err
}
