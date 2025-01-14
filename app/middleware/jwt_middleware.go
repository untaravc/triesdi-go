package middleware

import (
	"net/http"
	"strings"
	"triesdi/app/utils"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// If Content Type is empty or not application/json
		if c.ContentType() != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported Media Type"})
			c.Abort()
			return
		}

		// If There is no Bearer String
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := utils.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// You can set the claims to the context if needed
		c.Set("email", claims.Email)

		c.Next()
	}
}
