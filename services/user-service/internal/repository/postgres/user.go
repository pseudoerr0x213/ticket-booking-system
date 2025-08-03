package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error { return nil }

func (r *UserRepository) Delete(ctx context.Context, userID uuid.UUID) error { return nil }

func (r *UserRepository) Exists(ctx context.Context, userID uuid.UUID) (bool, error) {
	return false, nil
}
