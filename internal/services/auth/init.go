package auth

import (
	"bumimedika-backend/internal/domain/dto"
	models "bumimedika-backend/internal/models"
	authRepo "bumimedika-backend/internal/repositories/auth"
	"context"
)

type IAuthService interface {
	Register(ctx context.Context, user models.User) (*dto.UserDto, error)
	Login(ctx context.Context, email, password string) (*dto.TokenResponse, error)
}

type authService struct {
	repo   authRepo.IAuthRepository
	jwtKey string
}

func NewAuthService(repo authRepo.IAuthRepository, jwtKey string) IAuthService {
	return &authService{repo, jwtKey}
}
