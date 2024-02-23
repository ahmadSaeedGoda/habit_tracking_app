package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"lasting-dynamics.com/habit_tracking_app/internal/api/auth"
)

// IsAuthenticated is a middleware for Authorization header parsing & token validation
// Sets the user ID in the request context
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract the token without the "Bearer" prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the user ID in the context
		c.Set("userId", claims.UserID)
		c.Next()
	}
}
