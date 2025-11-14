package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SidBroGG/dementia-api/internal/model"
	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("not found")

type userDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *userDB {
	return &userDB{db: db}
}

func (r *userDB) Create(ctx context.Context, user *model.User) error {
	q := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRowxContext(ctx, q, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt)
}

func (r *userDB) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	q := `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`
	u := &model.User{}
	if err := r.db.GetContext(ctx, u, q, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return u, nil
}
