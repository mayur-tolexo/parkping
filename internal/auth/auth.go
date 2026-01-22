package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTestToken returns a JWT for testing
func GenerateTestToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Use the same secret as your middleware expects
	return token.SignedString([]byte("supersecret"))
}
