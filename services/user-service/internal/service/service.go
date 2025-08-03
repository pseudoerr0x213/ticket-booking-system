package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/repository"
)

type UserServicer interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*domain.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, req UpdateUserRequest) (*domain.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	ChangePassword(ctx context.Context, userID uuid.UUID, oldPassword, newPassword string) error
}

type AuthServicer interface {
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
	Logout(ctx context.Context, sessionToken string) error
	RefreshSession(ctx context.Context, sessionToken string) (*RefreshResponse, error)
	ValidateSession(ctx context.Context, sessionToken string) (*domain.User, error)
	LogoutAllSessions(ctx context.Context, userID uuid.UUID) error
}

type Services struct {
	User UserServicer
	Auth AuthServicer
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		User: NewUserService(repo),
		Auth: NewAuthService(repo),
	}
}

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required,min=1"`
	LastName  string `json:"last_name" validate:"required,min=1"`
}

type UpdateUserRequest struct {
	Email     *string `json:"email,omitempty" validate:"omitempty,email"`
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=1"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=1"`
}

type LoginResponse struct {
	User         *domain.User `json:"user"`
	SessionToken string       `json:"session_token"`
	ExpiresAt    time.Time    `json:"expires_at"`
}

type RefreshResponse struct {
	SessionToken string    `json:"session_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
