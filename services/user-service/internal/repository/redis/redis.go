package redis

import (
	"context"
	"time"

	"github.com/pseudoerr0x213/ticket-booking-system/user-service/internal/domain"
)

type UserCache interface {
	Set(ctx context.Context, key string, user *domain.User, ttl time.Duration) error
	Get(ctx context.Context, key string) (*domain.User, error)
}
