package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	IssueToken(ctx context.Context, userId int64) (token string, expiresAt time.Time, err error)
}

type JWTAuth struct {
	signingKey []byte
	ttl        time.Duration
}

func NewJWTAuth(signingKey []byte, ttl time.Duration) *JWTAuth {
	return &JWTAuth{signingKey: signingKey, ttl: ttl}
}

func (a *JWTAuth) IssueToken(ctx context.Context, userId int64) (string, time.Time, error) {
	expires := time.Now().Add(a.ttl)
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": jwt.NewNumericDate(expires),
		"iat": jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, expires, nil
}
