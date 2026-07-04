package database

import (
	"database/sql"
	"fmt"

	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/config"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func New(host, port, user, password, dbname string) (*Database, error) {
	config := config.NewEnvironmentConfig()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
