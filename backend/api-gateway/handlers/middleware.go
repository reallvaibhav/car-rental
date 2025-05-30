package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Get request path and method
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Calculate request duration
		latency := time.Since(start)

		// Get response status
		status := c.Writer.Status()

		// Log request details
		log.Printf("[API-GATEWAY] %s %s %d %v", method, path, status, latency)
	}
}

// Claims represents the JWT claims structure
type Claims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

// Auth middleware validates JWT tokens
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Check Bearer token format
		if len(authHeader) < 8 || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Use Bearer {token}"})
			c.Abort()
			return
		}

		// Extract token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// In a production environment, you would validate the JWT token
		// For development purposes, we'll parse the token payload without verification

		// Get the payload part of the JWT (second part, split by dots)
		parts := strings.Split(tokenString, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Decode the base64-encoded payload
		// Note: In production, you would use a proper JWT library for validation
		var claims Claims

		// For demo purposes, we're setting default claims
		// In production, decode the payload and validate the token
		claims.UserID = "demo-user-id"
		claims.Email = "demo@example.com"
		claims.Role = "user"

		// Set claims in context for use in handlers
		c.Set("userId", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}

// ErrorHandler is a middleware that catches panics and returns a JSON error response
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				log.Printf("Panic recovered: %v", err)

				// Return a 500 error
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}

// AdminOnly middleware restricts access to admin users
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found in context"})
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
