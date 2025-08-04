package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/service"
)

type UpdateUserRequest struct {
	Email     *string `json:"email,omitempty" binding:"omitempty,email"`
	FirstName *string `json:"first_name,omitempty" binding:"omitempty,min=1"`
	LastName  *string `json:"last_name,omitempty" binding:"omitempty,min=1"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=6"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

func (h *Handler) GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		ErrorResponse(c, http.StatusInternalServerError, "User not found in context", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func (h *Handler) GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	user, err := h.src.User.GetUser(c.Request.Context(), userID)
	if err != nil {
		ErrorResponse(c, http.StatusNotFound, "User not found", err)
		return
	}

	SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func (h *Handler) UpdateCurrentUser(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ValidationErrorResponse(c, err)
		return
	}

	user, exists := c.Get("user")
	if !exists {
		ErrorResponse(c, http.StatusInternalServerError, "User not found in context", nil)
		return
	}

	currentUser := user.(*domain.User)

	serviceReq := service.UpdateUserRequest{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	updatedUser, err := h.src.User.UpdateUser(c.Request.Context(), currentUser.ID, serviceReq)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Update failed", err)
		return
	}

	SuccessResponse(c, http.StatusOK, "User updated successfully", updatedUser)
}

func (h *Handler) DeleteCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		ErrorResponse(c, http.StatusInternalServerError, "User not found in context", nil)
		return
	}

	currentUser := user.(*domain.User)

	err := h.src.User.DeleteUser(c.Request.Context(), currentUser.ID)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Delete failed", err)
		return
	}

	SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ValidationErrorResponse(c, err)
		return
	}

	user, exists := c.Get("user")
	if !exists {
		ErrorResponse(c, http.StatusInternalServerError, "User not found in context", nil)
		return
	}

	currentUser := user.(*domain.User)

	err := h.src.User.ChangePassword(c.Request.Context(), currentUser.ID, req.OldPassword, req.NewPassword)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Password change failed", err)
		return
	}
}
