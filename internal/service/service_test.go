package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/SidBroGG/dementia-api/internal/model"
	"github.com/SidBroGG/dementia-api/internal/service"
	"github.com/SidBroGG/dementia-api/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// mocks

type mockUserStore struct {
	CreateFn     func(ctx context.Context, u *model.User) error
	GetByEmailFn func(ctx context.Context, email string) (*model.User, error)
}

func (m *mockUserStore) Create(ctx context.Context, u *model.User) error {
	return m.CreateFn(ctx, u)
}

func (m *mockUserStore) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return m.GetByEmailFn(ctx, email)
}

type mockAuth struct {
	IssueTokenFn func(ctx context.Context, id int64) (string, time.Time, error)
}

func (m *mockAuth) IssueToken(ctx context.Context, id int64) (string, time.Time, error) {
	return m.IssueTokenFn(ctx, id)
}

// tests

func TestRegister_Success(t *testing.T) {
	users := &mockUserStore{
		CreateFn: func(ctx context.Context, u *model.User) error {
			assert.NotEmpty(t, u.PasswordHash)
			assert.Equal(t, "test@test.com", u.Email)
			return nil
		},
	}

	s := service.NewService(store.Store{
		Users: users,
		Tasks: nil,
	}, nil)

	err := s.Register(context.Background(), model.AuthRequest{
		Email:    "	TeST@Test.com  ",
		Password: "testpassword",
	})

	assert.NoError(t, err)
}

func TestLogin_Success(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)

	users := &mockUserStore{
		GetByEmailFn: func(ctx context.Context, email string) (*model.User, error) {
			return &model.User{
				ID:           1,
				Email:        "test@mail.com",
				PasswordHash: string(hashed),
			}, nil
		},
	}

	auth := &mockAuth{
		IssueTokenFn: func(ctx context.Context, id int64) (string, time.Time, error) {
			return "token123", time.Time{}, nil
		},
	}

	s := service.NewService(store.Store{
		Users: users,
	}, auth)

	resp, err := s.Login(context.Background(), model.AuthRequest{
		Email:    "test@mail.com",
		Password: "testpassword",
	})

	require.NoError(t, err)
	assert.Equal(t, "token123", resp.Token)
}
