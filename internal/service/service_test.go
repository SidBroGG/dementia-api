package service

import (
	"context"
	"testing"
	"time"

	"github.com/SidBroGG/dementia-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type mockStore struct {
	mock.Mock
	lastCreated *model.User
}

func (m *mockStore) Create(ctx context.Context, u *model.User) error {
	args := m.Called(ctx, u)

	m.lastCreated = u
	return args.Error(0)
}

func (m *mockStore) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	u := args.Get(0)
	if u == nil {
		return nil, args.Error(1)
	}

	return u.(*model.User), args.Error(1)
}

type mockAuth struct {
	mock.Mock
}

func (m *mockAuth) IssueToken(ctx context.Context, userID int64) (string, time.Time, error) {
	args := m.Called(ctx, userID)
	token := ""

	if args.Get(0) != nil {
		token = args.String(0)
	}

	var t time.Time
	if args.Get(1) != nil {
		if tt, ok := args.Get(1).(time.Time); ok {
			t = tt
		}
	}

	return token, t, args.Error(2)
}

func TestRegisterSuccess(t *testing.T) {
	ctx := context.Background()
	ms := new(mockStore)
	ma := new(mockAuth)

	svc := NewService(ms, ma)

	req := model.AuthRequest{
		Email:    "   SKIBIDI.TOilet@qwe.com ",
		Password: "passssss123@#ord",
	}

	ms.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

	err := svc.Register(ctx, req)
	assert.NoError(t, err)

	created := ms.lastCreated
	if assert.NotNil(t, created) {
		assert.Equal(t, "skibidi.toilet@qwe.com", created.Email)

		err = bcrypt.CompareHashAndPassword([]byte(created.PasswordHash), []byte(req.Password))
		assert.NoError(t, err)
	}

	ms.AssertExpectations(t)
}
