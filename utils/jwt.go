package utils

import (
	"gopay-clone/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Function to create JWT tokens with claims
func CreateToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	var secretKey = []byte(secret)

	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,    // Subject (user identifier)
		"email":     user.Email, // Subject (user identifier)
		"user_type": user.Type,
		"exp":       time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":       time.Now().Unix(),                // Issued at
	})
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// function to get id and username from jwt
func CLaimJwt(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return int(uint(claims["user_id"].(float64)))
}
