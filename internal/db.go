package internal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type Storage struct {
	db *sqlx.DB
}

// God forbid to do that in production!
type Transaction struct {
	ID          uuid.UUID
	Time        time.Time
	State       string
	Amount      float32
	Balance     float32
	Transaction string
}

const dbPrefix = "db"

func NewStorage(dbName, dbUser, ssl, pwd, host, port string) (*Storage, error) {
	if dbName == "" || dbUser == "" || ssl == "" || pwd == "" || host == "" || port == "" {
		return nil, fmt.Errorf("%s: some db initialisation arguments are missing", dbPrefix)
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

func (s *Storage) UpdateBalance(ctx context.Context, t Transaction) (uuid.UUID, error) {
	// Check the uniqueness of the transaction
	if err := s.transactionIDCheck(ctx, t.Transaction); err != nil {
		return uuid.Nil, err
	}
	// Get the last stored balance
	val, err := s.getBalance(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	// Check the balance and perform the necessary operations
	switch t.State {
	case "win":
		t.Balance = t.Amount + val
	case "lost":
		if result := val - t.Amount; result >= 0 {
			t.Balance = result
		} else {
			return uuid.Nil, fmt.Errorf("%s: the overal balance cannot be negative", dbPrefix)
		}
	default:
		// Here we can panic, but I'm too laze to handle it, so...err
		return uuid.Nil, fmt.Errorf("%s: unknown state: %s", dbPrefix, t.State)
	}

	query := `insert into user_balance(id, time, state, amount, balance, transaction)
			values(:id, :time, :state, :amount, :balance, :transaction)`
	_, err = s.db.NamedExecContext(ctx, query, t)
	if err != nil {
		return uuid.Nil, err
	}
	return t.ID, nil
}

func (s *Storage) transactionIDCheck(ctx context.Context, id string) error {
	// Fetch the last element form the table
	// query := "select * from user_balance where transaction=$1 order by time desc limit 1"
	query := "select * from user_balance where transaction=$1 limit 1"
	row := s.db.QueryRowxContext(ctx, query, id)
	var transaction Transaction
	// Check transaction. If true, then the transaction was already stored; return.
	if err := row.StructScan(&transaction); err == sql.ErrNoRows {
		return nil
	} else {
		return err
	}
}

func (s *Storage) getBalance(ctx context.Context) (float32, error) {
	query := "select balance from user_balance order by time desc limit 1"
	row := s.db.QueryRowContext(ctx, query)
	var balance float32
	if err := row.Scan(&balance); err != nil {
		return 0, err
	}
	return balance, nil
}
