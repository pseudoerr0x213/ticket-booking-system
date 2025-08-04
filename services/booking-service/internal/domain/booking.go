package domain

import (
	"time"

	"github.com/google/uuid"
)

type BookingStatus uint

type Booking struct {
	ID          uuid.UUID     `json:"id"`
	UserID      uuid.UUID     `json:"user_id"`
	EventID     uuid.UUID     `json:"event_id"`
	TicketCount int           `json:"ticket_count"`
	TotalAmount float64       `json:"total_amount"`
	Status      BookingStatus `json:"status"`
	ExpiresAt   *time.Time    `json:"expires_at,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

const (
	BookingPending BookingStatus = iota
	BookingConfirmed
	BookingCancelled
	BookingExpired
	BookingRefunded
)

type Ticket struct {
	ID        uuid.UUID `json:"id"`
	BookingID uuid.UUID `json:"booking_id"`
	SeatInfo  *string   `json:"seat_info,omitempty"`
	QRCode    string    `json:"qr_code"`
	CreatedAt time.Time `json:"created_at"`
}

type BookingRequest struct {
	EventID     uuid.UUID `json:"event_id" validate:"required"`
	TicketCount int       `json:"ticket_count" validate:"required,min=1,max=10"`
	UserID      uuid.UUID // This comes from authenticated context, not request body
}
