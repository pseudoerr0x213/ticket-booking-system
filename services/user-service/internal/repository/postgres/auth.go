package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateSession(ctx context.Context, session *domain.Session) error {
	return nil
}

func (r *AuthRepository) GetSessionByToken(ctx context.Context, sessionToken string) (*domain.Session, error) {
	return nil, nil
}

func (r *AuthRepository) GetUserBySession(ctx context.Context, sessionToken string) (*domain.User, error) {
	return nil, nil
}

func (r *AuthRepository) UpdateSession(ctx context.Context, session *domain.Session) error {
	return nil
}

func (r *AuthRepository) DeleteSession(ctx context.Context, sessionToken string) error { return nil }

func (r *AuthRepository) DeleteUserSessions(ctx context.Context, userID uuid.UUID) error { return nil }

func (r *AuthRepository) CleanupExpiredSessions(ctx context.Context) error { return nil }
