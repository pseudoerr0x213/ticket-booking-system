package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/repository"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	return nil, nil
}

func (s *AuthService) Logout(ctx context.Context, sessionToken string) error { return nil }

func (s *AuthService) RefreshSession(ctx context.Context, sessionToken string) (*RefreshResponse, error) {
	return nil, nil
}

func (s *AuthService) ValidateSession(ctx context.Context, sessionToken string) (*domain.User, error) {
	return nil, nil
}

func (s *AuthService) LogoutAllSessions(ctx context.Context, userID uuid.UUID) error { return nil }
