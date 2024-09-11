package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func ValidateJWTFromHeader(c *gin.Context) (*Claims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header missing")
	}

	tokenString := strings.Split(authHeader, "Bearer ")[1]
	return ValidateJWT(tokenString)
}

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
