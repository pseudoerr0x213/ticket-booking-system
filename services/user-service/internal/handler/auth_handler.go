package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/service"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name" binding:"required,min=1"`
	LastName  string `json:"last_name" binding:"required,min=1"`
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ValidationErrorResponse(c, err)
		return
	}

	response, err := h.src.Auth.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		ErrorResponse(c, http.StatusUnauthorized, "Login failed", err)
		return
	}

	SuccessResponse(c, http.StatusOK, "Login successful", response)
}

func (h *Handler) Signup(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ValidationErrorResponse(c, err)
		return
	}

	serviceReq := service.CreateUserRequest{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	user, err := h.src.User.CreateUser(c.Request.Context(), serviceReq)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Registration failed", err)
		return
	}

	SuccessResponse(c, http.StatusCreated, "User registered successfully", user)
}

func (h *Handler) Logout(c *gin.Context) {
	sessionToken, exists := c.Get("session_token")
	if !exists {
		ErrorResponse(c, http.StatusInternalServerError, "Session token not found", nil)
		return
	}

	err := h.src.Auth.Logout(c.Request.Context(), sessionToken.(string))
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Logout failed", err)
		return
	}

	SuccessResponse(c, http.StatusOK, "Logged out successfully", nil)
}

func (h *Handler) RefreshSession(c *gin.Context) {
	sessionToken, exists := c.Get("session_token")
	if !exists {
		ErrorResponse(c, http.StatusInternalServerError, "Session token not found", nil)
		return
	}

	response, err := h.src.Auth.RefreshSession(c.Request.Context(), sessionToken.(string))
	if err != nil {
		ErrorResponse(c, http.StatusUnauthorized, "Session refresh failed", err)
		return
	}

	SuccessResponse(c, http.StatusOK, "Session refreshed successfully", response)
}
