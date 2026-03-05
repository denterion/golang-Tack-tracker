package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewDB() (*pgx.Conn, error) {
	connStr := "postgres://postgres:123@localhost:5432/tasktracker"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return conn, nil
}
