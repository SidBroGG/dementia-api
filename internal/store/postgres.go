package store

import (
	"fmt"

	"github.com/SidBroGG/dementia-api/config"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg config.DB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("cant connect to db: %w", err)
	}

	return db, nil
}
