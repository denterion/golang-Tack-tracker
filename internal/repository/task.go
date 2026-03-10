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

// CreateTask godoc
// @Summary Создать задачу
// @Description Создает новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task data"
// @Success 201 {object} model.Task
// @Failure 400 {string} string
// @Router /tasks [post]
func (r *Repository) CreateTask(t *model.Task) error {
	sql := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRow(context.Background(), sql, t.Title, t.Description, t.Status).Scan(&t.ID, &t.CreatedAt)
}
// GetTasks godoc
// @Summary Получить список задач
// @Tags tasks
// @Produce json
// @Success 200 {array} model.Task
// @Router /tasks [get]
func (r *Repository) GetTasks() ([]*model.Task, error) {
	sql := `SELECT id, title, description, status, created_at FROM tasks ORDER BY id`
	rows, err := r.db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task

	for rows.Next() {
		t := &model.Task{}
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
// GetTask godoc
// @Summary Получить задачу
// @Tags tasks
// @Param id path int true "Task ID"
// @Produce json
// @Success 200 {object} model.Task
// @Router /tasks/{id} [get]
func (r *Repository) GetTaskByID(id int) (*model.Task, error) {
	sql := `SELECT id, title, description, status, created_at FROM tasks WHERE id=$1`

	var t model.Task

	err := r.db.QueryRow(context.Background(), sql, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
// UpdateTask godoc
// @Summary Обновить задачу
// @Tags tasks
// @Param id path int true "Task ID"
// @Param task body model.Task true "Task"
// @Produce json
// @Success 200 {object} model.Task
// @Router /tasks/{id} [put]
func (r *Repository) UpdateTask(t *model.Task) error {
	sql := `UPDATE tasks
	SET title=$1, description=$2, status=$3
	WHERE id=$4`
	_, err := r.db.Exec(context.Background(), sql, t.Title, t.Description, t.Status, t.ID)
	return err
}
// DeleteTask godoc
// @Summary Удалить задачу
// @Tags tasks
// @Param id path int true "Task ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Router /tasks/{id} [delete]
func (r *Repository) DeleteTask(id int) error {
	sql := `DELETE FROM tasks WHERE id=$1`
	_, err := r.db.Exec(context.Background(), sql, id)
	return err
}
