package helper

import (
	"log"
	"time"
	"github.com/Gylmynnn/go-gin-flutter-bloc/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "malas"))

func GenerateToken(username string) string {
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		log.Println("peringatan : gagal membuat token baru")
	}

	return token
}
