package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pseudoerr0x213/ticket-booking-system/booking-service/internal/domain"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *domain.Booking) error
	GetByID(ctx context.Context, bookingID uuid.UUID) (*domain.Booking, error)
	Update(ctx context.Context, booking *domain.Booking) error
	Delete(ctx context.Context, bookingID uuid.UUID) error

	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Booking, error)
	GetByEventID(ctx context.Context, eventID uuid.UUID, limit, offset int) ([]*domain.Booking, error)
	GetByStatus(ctx context.Context, status domain.BookingStatus, limit, offset int) ([]*domain.Booking, error)

	GetActiveBookingsByUser(ctx context.Context, userID uuid.UUID) ([]*domain.Booking, error)
	GetExpiredBookings(ctx context.Context, before time.Time) ([]*domain.Booking, error)
	GetBookingsByDateRange(ctx context.Context, from, to time.Time, limit, offset int) ([]*domain.Booking, error)

	Exists(ctx context.Context, bookingID uuid.UUID) (bool, error)
	CountByUserID(ctx context.Context, userID uuid.UUID) (int, error)
	CountByEventID(ctx context.Context, eventID uuid.UUID) (int, error)

	UpdateStatus(ctx context.Context, bookingIDs []uuid.UUID, status domain.BookingStatus) error
	GetBookingsByIDs(ctx context.Context, bookingIDs []uuid.UUID) ([]*domain.Booking, error)
}

type TicketRepository interface {
	Create(ctx context.Context, ticket *domain.Ticket) error
	CreateBatch(ctx context.Context, tickets []*domain.Ticket) error
	GetByID(ctx context.Context, ticketID uuid.UUID) (*domain.Ticket, error)
	Update(ctx context.Context, ticket *domain.Ticket) error

	GetByBookingID(ctx context.Context, bookingID uuid.UUID) ([]*domain.Ticket, error)
	GetByQRCode(ctx context.Context, qrCode string) (*domain.Ticket, error)

	ValidateTicket(ctx context.Context, qrCode string) (*domain.Ticket, error)
	Exists(ctx context.Context, ticketID uuid.UUID) (bool, error)
}

type Repository struct {
	Booking BookingRepository
	Ticket  TicketRepository
}
