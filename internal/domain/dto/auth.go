package dto

import (
	entities "bumimedika-backend/internal/domain/entities"
	"time"
)

type UserDto struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	RoleID    uint64    `json:"role_id"`
	Role      RoleDto   `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleDto struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func ToUserDto(u *entities.User) *UserDto {
	if u == nil {
		return nil
	}
	return &UserDto{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		Role:      ToRoleDto(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToRoleDto(r entities.Role) RoleDto {
	return RoleDto{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
	}
}
