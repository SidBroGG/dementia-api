package store

import "github.com/jmoiron/sqlx"

type Store struct {
	DB    *sqlx.DB
	Users UserRepo
	Tasks TaskRepo
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		DB:    db,
		Users: NewUserDB(db),
		Tasks: nil, // TODO
	}
}
