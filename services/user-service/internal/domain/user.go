package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	Status       UserStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserValidator struct{}

type UserStatus int

const (
	Admin UserStatus = iota
	Customer
	Organizer
)
