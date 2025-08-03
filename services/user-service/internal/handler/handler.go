package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/service"
)

type Handler struct {
	src *service.Services
	mw  *Middleware
}

func NewHandler(src *service.Services, mw *Middleware) *Handler {
	return &Handler{
		src: src,
		mw:  mw,
	}
}

func (h *Handler) SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/health", h.HealthCheck)
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", h.Login)
			auth.POST("/signup", h.Signup)
			auth.POST("/logout", h.mw.Auth(), h.Logout)
			auth.POST("/refresh", h.mw.Auth(), h.RefreshSession)
		}

		// User routes
		users := v1.Group("/users")
		{
			users.GET("/me", h.mw.Auth(), h.GetCurrentUser)
			users.PUT("/me", h.mw.Auth(), h.UpdateCurrentUser)
			users.DELETE("/me", h.mw.Auth(), h.DeleteCurrentUser)
			users.POST("/change-password", h.mw.Auth(), h.ChangePassword)

			// Admin routes
			users.GET("/:id", h.mw.Auth(), h.GetUser)
		}
	}

	return r
}

func (h *Handler) HealthCheck(c *gin.Context) {
	SuccessResponse(c, 200, "Service is healthy", gin.H{
		"status":  "ok",
		"service": "user-service",
	})
}
