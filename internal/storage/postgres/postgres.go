package postgres

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/hahaclassic/elysium/config"
	"github.com/hahaclassic/elysium/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, config *config.PostgresConfig) (*Storage, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		config.User, config.Password, net.JoinHostPort(config.Host, config.Port), config.DB, config.SSLMode)

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, errors.Join(storage.ErrStorageConnection, err)
	}

	if db.Ping(ctx) != nil {
		return nil, errors.Join(storage.ErrStorageConnection, err)
	}

	return &Storage{db}, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}
