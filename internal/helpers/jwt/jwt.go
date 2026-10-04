package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/magomed066/auth-go-app/internal/env"
)

type Claims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

func CreateToken(userID uint) (string, error) {
	secret := env.GetEnvs().JWT_ACCESS_TOKEN
	secretExpiresIn := env.GetEnvs().JWT_ACCESS_EXPIRES_IN

 	expiresIn := time.Duration(secretExpiresIn) * time.Hour
	
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}