package repository

import (
	"database/sql" 
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.ErrorLog ) error {

	query := `
		INSERT INTO users(name, email)
		VALUES($1, $2)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		user.Description,
		user.ID,
	).Scan(&user.ID)
}
