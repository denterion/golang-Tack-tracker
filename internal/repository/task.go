package repository

import (
	"context"
	"task-tracker/internal/model"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTask(t *model.Task) error {
	sql := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRow(context.Background(), sql, t.Title, t.Description, t.Status).Scan(&t.ID, &t.CreatedAt)
}
