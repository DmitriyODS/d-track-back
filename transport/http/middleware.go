package http

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
	"net/http"
)

// JWTAuth - промежуточное ПО для проверки нашего токена (аутентификации)
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get(global.HeaderAuthenticationKey)
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		claims := &global.JwtClaims{}
		token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.JwtSecretKey), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		c.Set(global.JwtClaimsCtxKey, claims)

		c.Next()
	}
}

// RecoveryPanic - обрабатывает любую панику во время работы сервера
func RecoveryPanic(c *gin.Context, err any) {
	log.Println("RecoveryPanic err:", err)
	c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
}
