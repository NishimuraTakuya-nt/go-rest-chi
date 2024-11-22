package usecase

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/service"
)

type AuthUsecase interface {
	Login(ctx context.Context, userID string, roles []string) (string, error)
	Authenticate(ctx context.Context, tokenString string) (*model.User, error)
}

type authUsecase struct {
	tokenService service.TokenService
}

func NewAuthUsecase(tokenService service.TokenService) AuthUsecase {
	return &authUsecase{
		tokenService: tokenService,
	}
}

func (uc *authUsecase) Login(ctx context.Context, userID string, roles []string) (string, error) {
	return uc.tokenService.GenerateToken(ctx, userID, roles)
}

func (uc *authUsecase) Authenticate(ctx context.Context, tokenString string) (*model.User, error) {
	claims, err := uc.tokenService.ValidateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    claims.UserID,
		Roles: claims.Roles,
	}, nil
}
