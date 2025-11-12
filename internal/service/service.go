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

type Service struct {
	store store.Store
	auth  auth.Auth
}

func NewService(store store.Store, auth auth.Auth) *Service {
	return &Service{store: store, auth: auth}
}

func (s *Service) Register(ctx context.Context, req model.AuthRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	user := &model.User{
		Email:        strings.ToLower(strings.TrimSpace(req.Email)),
		PasswordHash: string(hashed),
	}

	if err := s.store.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(ctx context.Context, req model.AuthRequest) (*model.LoginResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	user, err := s.store.GetByEmail(ctx, email)
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
