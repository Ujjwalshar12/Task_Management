package repository

import (
	"database/sql"
	"task_management/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *model.User) error {

	_, err := r.DB.Exec(`
		INSERT INTO users (id, email, password, role)
		VALUES ($1, $2, $3, $4)
	`,
		user.ID,
		user.Email,
		user.Password,
		user.Role,
	)

	return err
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {

	row := r.DB.QueryRow(`
		SELECT id, email, password, role 
		FROM users WHERE email=$1
	`, email)

	var u model.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role)

	return &u, err
}
