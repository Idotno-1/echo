package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"idotno.fr/echo/models"
)

// var jwtKey = []byte(GenerateJWTKey())
var jwtKey = []byte("aaaaaaaaaaa")

type Claims struct {
	UserID   uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWTKey() string {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal("Error generating random key:", err)
	}

	return base64.StdEncoding.EncodeToString(key)
}

func GenerateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
