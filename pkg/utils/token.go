package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(UserID string, tokenType string, exp int, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    UserID,
		"token_type": tokenType,
		"exp":        time.Now().Add(time.Duration(exp) * time.Second).Unix(),
		"iat":        time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return result, nil
}
