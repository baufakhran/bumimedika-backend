package auth

import (
	"context"
	"errors"
	"time"

	"bumimedika-backend/internal/domain/dto"
	"bumimedika-backend/internal/domain/entities"
	models "bumimedika-backend/internal/models"
	"bumimedika-backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// Register creates a new user account
func (s *authService) Register(ctx context.Context, payload models.User) (*dto.UserDto, error) {
	existing, err := s.repo.FindByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Email:     payload.Email,
		Password:  string(hashed),
		Name:      payload.Name,
		RoleID:    payload.RoleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return dto.ToUserDto(user), nil
}

// Login verifies credentials and returns JWT
func (s *authService) Login(ctx context.Context, email, password string) (*dto.TokenResponse, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(s.jwtKey, map[string]interface{}{
		"user_id":   user.ID,
		"user_name": user.Name,
	})
	if err != nil {
		return nil, err
	}

	result := &dto.TokenResponse{Token: token}

	return result, nil
}
