package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/service"
)

type Middleware struct {
	authService service.AuthService
}

func NewMiddleware(authService service.AuthService) *Middleware {
	return &Middleware{
		authService: authService,
	}
}

func (m *Middleware) CORS() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

func (m *Middleware) Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			ErrorResponse(c, http.StatusUnauthorized, "Authorization header required", nil)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header format", nil)
			c.Abort()
			return
		}

		token := parts[1]
		user, err := m.authService.ValidateSession(c.Request.Context(), token)
		if err != nil {
			ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired session", err)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("session_token", token)
		c.Next()
	})
}

/* func (m *Middleware) RateLimit() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Next()
	})
}
*/
