package internal

import (
	"context"

	// "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Storage .
type Storage struct {
}

// NewStorage .
func NewStorage(cfg string) (*Storage, error) {
	// db, err := sqlx.Connect("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s", )
	return &Storage{}, nil
}

// UpdateBalance .
func (s *Storage) UpdateBalance(ctx context.Context, state string, amount float32, transactionID string) error {
	return nil
}
