package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(connectionString string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
