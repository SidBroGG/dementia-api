package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/SidBroGG/dementia-api/internal/auth"
	"github.com/SidBroGG/dementia-api/internal/model"
	"github.com/SidBroGG/dementia-api/internal/store"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCreds = errors.New("invalid credentials")

type AuthService struct {
	users store.UserRepo
	auth  auth.Auth
}

func NewAuthService(users store.UserRepo, auth auth.Auth) *AuthService {
	return &AuthService{users: users, auth: auth}
}

func (s *AuthService) Register(ctx context.Context, req model.AuthRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	user := &model.User{
		Email:        strings.ToLower(strings.TrimSpace(req.Email)),
		PasswordHash: string(hashed),
	}

	if err := s.users.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, req model.AuthRequest) (*model.LoginResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	user, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return nil, ErrInvalidCreds
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCreds
	}

	token, _, err := s.auth.IssueToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{Token: token}, nil
}
