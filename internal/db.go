package internal

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(dbName, dbUser, ssl, pwd, host, port string) (*Storage, error) {
	if dbName == "" || dbUser == "" || ssl == "" || pwd == "" || host == "" || port == "" {
		return nil, errors.New("error: some db initialisation arguments are missing")
	}
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		host, port, pwd, dbUser, dbName, ssl)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(2)
	return &Storage{db: db}, nil
}

func (s *Storage) UpdateBalance(ctx context.Context, state string, amount float64, transactionID string) (uuid.UUID, error) {
	if err := s.transactionIDCheck(ctx, transactionID); err != nil {
		return uuid.Nil, err
	}

	// Get the last value
	val, err := s.getTheLastStoredAmount()
	if err != nil {
		return uuid.Nil, err
	}

	query := "insert into user_balance(id, amount, transaction) values(default, $1, $2)"
	_, err = s.db.ExecContext(ctx, query, val, transactionID)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.UUID{}, nil
}

func (s *Storage) transactionIDCheck(ctx context.Context, id string) error {
	query := "select * from user_balance where transaction=$1"
	rows, err := s.db.QueryxContext(ctx, query, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	trans := ""
	// Extract transactionID, if it exists
	for rows.Next() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := rows.Scan(&trans); err != nil {
				return err
			}
		}
	}
	// Check transaction
	if trans != "" {
		return fmt.Errorf("error: transaction with id %s was already stored", trans)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (s *Storage) getTheLastStoredAmount() (float64, error) {
	// query := `select * from user_balance where user_name=$1 order by start_time asc`
	return 0, nil
}
