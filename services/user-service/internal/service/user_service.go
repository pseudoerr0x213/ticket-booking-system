package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*domain.User, error) {
	return nil, nil
}

func (s *UserService) GetUser(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userID uuid.UUID, req UpdateUserRequest) (*domain.User, error) {
	return nil, nil
}

func (s *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) error { return nil }

func (s *UserService) ChangePassword(ctx context.Context, userID uuid.UUID, oldPassword, newPassword string) error {
	return nil
}
