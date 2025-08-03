package cache

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	client redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		client: *client,
	}
}

func (r *CacheRepository) SetUser(ctx context.Context, user *domain.User, ttl time.Duration) error {
	return nil
}

func (r *CacheRepository) GetUser(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (r *CacheRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error { return nil }

func (r *CacheRepository) SetSession(ctx context.Context, session *domain.Session, ttl time.Duration) error {
	return nil
}

func (r *CacheRepository) GetSession(ctx context.Context, sessionToken string) (*domain.Session, error) {
	return nil, nil
}

func (r *CacheRepository) DeleteSession(ctx context.Context, sessionToken string) error { return nil }
