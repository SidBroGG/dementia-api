package store

import (
	"context"

	"github.com/SidBroGG/dementia-api/internal/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type TaskRepo interface {
	Create(ctx context.Context, user *model.Task) error
	GetByID(ctx context.Context, id int64) (*model.Task, error)
	List(ctx context.Context, userID int64, page int, pageSize int) ([]*model.Task, error)
	Update(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}
