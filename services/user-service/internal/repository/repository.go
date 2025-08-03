package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/repository/postgres"
	cache "github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/repository/redis"
	"github.com/redis/go-redis/v9"
)

type UserRepository interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, userID uuid.UUID) error
	Exists(ctx context.Context, userID uuid.UUID) (bool, error)
}

type AuthRepository interface {
	CreateSession(ctx context.Context, session *domain.Session) error
	GetSessionByToken(ctx context.Context, sessionToken string) (*domain.Session, error)
	GetUserBySession(ctx context.Context, sessionToken string) (*domain.User, error)
	UpdateSession(ctx context.Context, session *domain.Session) error
	DeleteSession(ctx context.Context, sessionToken string) error
	DeleteUserSessions(ctx context.Context, userID uuid.UUID) error
	CleanupExpiredSessions(ctx context.Context) error
}

type CacheRepository interface {
	SetUser(ctx context.Context, user *domain.User, ttl time.Duration) error
	GetUser(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	SetSession(ctx context.Context, session *domain.Session, ttl time.Duration) error
	GetSession(ctx context.Context, sessionToken string) (*domain.Session, error)
	DeleteSession(ctx context.Context, sessionToken string) error
}

type Repository struct {
	User  UserRepository
	Auth  AuthRepository
	Cache CacheRepository
}

func NewRepository(db *sql.DB, redisClient *redis.Client) *Repository {
	return &Repository{
		User:  postgres.NewUserRepository(db),
		Auth:  postgres.NewAuthRepository(db),
		Cache: cache.NewCacheRepository(redisClient),
	}
}
