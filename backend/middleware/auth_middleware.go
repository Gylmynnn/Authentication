package middleware

import (
	"net/http"
	"strings"

	"github.com/Gylmynnn/go-gin-flutter-bloc/config"
	"github.com/Gylmynnn/go-gin-flutter-bloc/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "malas"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		strToken := c.GetHeader("Authorization")

		if strToken == "" {
			c.JSON(http.StatusUnauthorized, helper.FailedResponse{
				Success: false,
				Status:  http.StatusUnauthorized,
				Message: "error : token tidak boleh kosong",
			})
			c.Abort()
			return
		}

		strToken = strings.TrimPrefix(strToken, "Bearer ")
		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(strToken, claims, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, helper.FailedResponse{
				Success: false,
				Status:  http.StatusUnauthorized,
				Message: "error : token tidak valid",
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Subject)
		c.Next()
	}
}
