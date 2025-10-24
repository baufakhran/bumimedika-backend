package config

import (
	"bumimedika-backend/internal/repositories/auth"
	"bumimedika-backend/internal/repositories/customer"
	"bumimedika-backend/internal/repositories/product"
	"bumimedika-backend/pkg/database"
	"bumimedika-backend/pkg/redis"

	rdsClient "github.com/redis/go-redis/v9"
)

type Repository struct {
	AuthRepository     auth.IAuthRepository
	ProductRepository  product.IProductRepository
	CustumerRepository customer.ICustumerRepository
	Rds                *rdsClient.Client
}

func InitRepository(cfg *Config) *Repository {
	rdb := redis.NewRedisClient(cfg.RedisHost, cfg.RedisPass, 0)
	db := database.NewMySQL(cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName)
	return &Repository{
		AuthRepository:     auth.InitAuthRepository(db),
		ProductRepository:  product.InitProductRepository(db),
		CustumerRepository: customer.InitCustomerRepository(db),
		Rds:                rdb,
	}
}
