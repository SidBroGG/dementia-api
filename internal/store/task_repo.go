package store

import (
	"context"
	"database/sql"

	"github.com/SidBroGG/dementia-api/internal/model"
	"github.com/jmoiron/sqlx"
)

type taskDB struct {
	db *sqlx.DB
}

func NewTaskDB(db *sqlx.DB) *taskDB {
	return &taskDB{db: db}
}

func (r *taskDB) Create(ctx context.Context, task *model.Task) error {
	q := `INSERT INTO tasks (user_id, title, description) 
		  VALUES ($1, $2, $3) 
		  RETURNING id, created_at, updated_at`
	return r.db.QueryRowxContext(ctx, q, task.UserID, task.Title, task.Description).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *taskDB) GetByID(ctx context.Context, id int64) (*model.Task, error) {
	q := `SELECT user_id, title, description, created_at, updated_at
		  FROM tasks WHERE id = $1`
	t := &model.Task{}

	if err := r.db.GetContext(ctx, t, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return t, nil
}

func (r *taskDB) Update(ctx context.Context, task *model.Task) error {
	q := `UPDATE tasks
		  SET title = $1, description = $2, updated_at = NOW()
		  WHERE id = $3
		  RETURNING updated_at`
	return r.db.QueryRowxContext(ctx, q, task.Title, task.Description, task.ID).Scan(&task.UpdatedAt)
}

func (r *taskDB) Delete(ctx context.Context, id int64) error {
	q := `DELETE FROM tasks WHERE id = $1`
	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *taskDB) List(ctx context.Context, userID int64, page int, pageSize int) ([]*model.Task, error) {
	offset := (page - 1) * pageSize

	q := `SELECT id, user_id, title, description, created_at, updated_at
		  FROM tasks
		  WHERE	user_id = $1
		  ORDER BY created_at DESC
		  LIMIT $2 OFFSET $3`

	tasks := []*model.Task{}
	if err := r.db.SelectContext(ctx, &tasks, q, userID, pageSize, offset); err != nil {
		return nil, err
	}

	return tasks, nil
}
