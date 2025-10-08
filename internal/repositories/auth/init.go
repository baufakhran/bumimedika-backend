package auth

import (
	"context"

	entities "bumimedika-backend/internal/domain/entities"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
}

type authRepository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db}
}
