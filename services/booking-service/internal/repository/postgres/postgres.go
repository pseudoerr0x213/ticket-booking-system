package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pseudoerr0x213/ticket-booking-system/booking-service/internal/config"
)

func NewPostgresDB(cfg config.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
