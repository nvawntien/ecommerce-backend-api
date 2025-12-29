package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(UserID string, userRole string, exp int, secret string) (string, error) {
    claims := jwt.RegisteredClaims{
        Subject:  UserID,               
        Audience: jwt.ClaimStrings{userRole}, 
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exp) * time.Second)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}