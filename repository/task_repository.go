package repository

import (
	"database/sql"
	"task_management/model"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) Create(task *model.Task) error {
	_, err := r.DB.Exec(`
		INSERT INTO tasks VALUES ($1,$2,$3,$4,$5,$6,$7)`, task.ID, task.Title, task.Description, task.Status,
		task.UserID, task.CreatedAt, task.UpdatedAt)
	return err
}

func (r *TaskRepository) GetByID(id, userID, role string) (*model.Task, error) {
	query := "SELECT * FROM tasks WHERE id=$1"
	args := []any{id}

	if role != "admin" {
		query += " AND user_id=$2"
		args = append(args, userID)
	}

	row := r.DB.QueryRow(query, args...)

	var t model.Task
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Status,
		&t.UserID, &t.CreatedAt, &t.UpdatedAt)

	return &t, err
}

func (r *TaskRepository) List(userID, role string) ([]model.Task, error) {
	query := "SELECT * FROM tasks"
	args := []any{}

	if role != "admin" {
		query += " WHERE user_id=$1"
		args = append(args, userID)
	}

	rows, _ := r.DB.Query(query, args...)
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status,
			&t.UserID, &t.CreatedAt, &t.UpdatedAt)
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepository) Delete(id, userID, role string) error {
	query := "DELETE FROM tasks WHERE id=$1"
	args := []any{id}

	if role != "admin" {
		query += " AND user_id=$2"
		args = append(args, userID)
	}

	_, err := r.DB.Exec(query, args...)
	return err
}

func (r *TaskRepository) AutoComplete(id string) {
	r.DB.Exec(`
		UPDATE tasks 
		SET status='completed', updated_at=NOW()
		WHERE id=$1 AND status IN ('pending','in_progress')
	`, id)
}
