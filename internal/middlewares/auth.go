package middlewares

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsExpiredRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil || refreshToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    10402,
				"message": "Không tìm thấy Refresh Token, vui lòng đăng nhập lại",
			})
			return
		}
	
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(global.Config.JWT.RefreshSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    10402,
				"message": "Refresh Token không hợp lệ hoặc đã hết hạn",
			})
			return
		}

		userID := claims.Subject
		
		var userRole string
		if len(claims.Audience) > 0 {
			userRole = claims.Audience[0]
		}

		c.Set("user_id", userID)
		c.Set("role", userRole)
		
		c.Next()
	}
}