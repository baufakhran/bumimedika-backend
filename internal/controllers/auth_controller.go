package controllers

import (
	models "bumimedika-backend/internal/models"
	authService "bumimedika-backend/internal/services/auth"
	"bumimedika-backend/pkg/response"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type AuthController struct {
	service authService.IAuthService
}

func InitAuthController(s authService.IAuthService) IAuthController {
	return &AuthController{s}
}

// POST /api/register
func (h *AuthController) Register(c *gin.Context) {
	var payload models.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(context.Background(), payload)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: user}
		response.Success(c)
	}
}

// POST /api/login
func (h *AuthController) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(context.Background(), req.Email, req.Password)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: token}
		response.Success(c)
	}
}
