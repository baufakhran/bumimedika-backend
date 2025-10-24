package config

import (
	"bumimedika-backend/internal/services/auth"
	"bumimedika-backend/internal/services/customer"
	"bumimedika-backend/internal/services/product"
)

type Usecase struct {
	// Add usecase interfaces here
	AuthUsecase     auth.IAuthService
	ProductUsecase  product.IProductService
	CustomerUsecase customer.ICustumerService
}

func InitUsecase(repo *Repository, cfg *Config) *Usecase {
	return &Usecase{
		AuthUsecase:     auth.NewAuthService(repo.AuthRepository, cfg.JWTKey),
		ProductUsecase:  product.NewProductService(repo.ProductRepository),
		CustomerUsecase: customer.InitCustumerService(repo.CustumerRepository),
	}
}
