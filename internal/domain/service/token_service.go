package service

import (
	"context"
	"errors"
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-chi/config"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(ctx context.Context, userID string, roles []string) (string, error)
	ValidateToken(ctx context.Context, tokenString string) (*model.Claims, error)
}

type tokenService struct {
	cfg *config.AppConfig
}

func NewTokenService(cfg *config.AppConfig) TokenService {
	return &tokenService{
		cfg: cfg,
	}
}

func (s *tokenService) GenerateToken(_ context.Context, userID string, roles []string) (string, error) {
	claims := &model.Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecretKey))
}

func (s *tokenService) ValidateToken(_ context.Context, tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(_ *jwt.Token) (any, error) {
		return []byte(s.cfg.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}