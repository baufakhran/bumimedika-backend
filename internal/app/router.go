package app

import (
	"bumimedika-backend/internal/app/config"
	"bumimedika-backend/internal/middleware"

	redis "github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

type Routes struct {
	controller config.Controller
}

type IRoutes interface {
	registerRoutes(cfg *config.Config, rdb *redis.Client)
}

// RegisterRoutes is used to register url routes API
func initRoutes(controller config.Controller) IRoutes {
	return &Routes{
		controller: controller,
	}
}

func (r *Routes) registerRoutes(cfg *config.Config, rdb *redis.Client) {
	// Initialize Gin router
	router = gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Public routes auth
	api := router.Group("/auth")
	api.POST("/login", r.controller.AuthController.Login)
	api.POST("/register", r.controller.AuthController.Register)

	productRoutes := router.Group("/product")
	// productRoutes.Use(middleware.AuthMiddleware(cfg.JWTKey, rdb))
	productRoutes.GET("", r.controller.ProductController.GetList)
	productRoutes.POST("", r.controller.ProductController.Manage)
	productRoutes.GET("/:id", r.controller.ProductController.GetByID)
	productRoutes.DELETE("/:id", r.controller.ProductController.DeleteByID)

	customerRoutes := router.Group("/customer")
	// customerRoutes.Use(middleware.AuthMiddleware(cfg.JWTKey, rdb))
	customerRoutes.GET("", r.controller.CustomerController.GetList)
	customerRoutes.POST("", r.controller.CustomerController.Manage)
	customerRoutes.GET("/:id", r.controller.CustomerController.GetByID)
	customerRoutes.DELETE("/:id", r.controller.CustomerController.DeleteByID)
}
