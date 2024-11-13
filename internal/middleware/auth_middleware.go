package middlewares

import (
	"net/http"
	"strings"
	"golang-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c * gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"token expirado"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {

			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "formato de token inválido"})
			c.Abort()
			return 
		}

		token := tokenParts[1]

		claims, err := utils.ValidateToken(token)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "token inválido"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}