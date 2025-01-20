package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"
	"triesdi/app/utils"

	"github.com/gin-gonic/gin"
)

func processTimeout(h gin.HandlerFunc, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		processDone := make(chan bool)
		go func() {
			h(c)
			processDone <- true
		}()

		select {
		case <-ctx.Done():
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "timeout wkwkwk"})
			c.Abort()
		case <-processDone:
		}
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return processTimeout(func(c *gin.Context) {
		
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Only check Content-Type for PATCH and POST methods exclude FILE
		if (c.Request.Method == http.MethodPatch || c.Request.Method == http.MethodPost) && !strings.Contains(c.ContentType(), "multipart/form-data") {

			// If Content Type is empty or not application/json and url not v1/file
			if c.ContentType() != "application/json" && !strings.Contains(c.Request.URL.Path, "/v1/file") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported Media Type"})
				c.Abort()
				return
			}
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
		c.Set("id", claims.ID)

		c.Next()
	}, 2*time.Second)
}
