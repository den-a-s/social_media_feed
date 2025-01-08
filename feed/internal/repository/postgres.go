package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db_connected_string := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, os.Getenv("DB_PASSWORD"), cfg.SSLMode)

	db, err := sqlx.Connect("postgres", db_connected_string)
	if err != nil {
		panic(fmt.Sprintf("DB is not connected %s", err))
	}

	return db, nil
}
