package product

import (
	"bumimedika-backend/internal/domain/dto"
	repository "bumimedika-backend/internal/repositories/product"
)

type IProductService interface {
	Manage(product *dto.Product) error
	GetByID(id uint64) (*dto.Product, error)
	GetList(name, sortField, sortOrder string, page, limit int) ([]dto.Product, int64, error)
}

type productService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) IProductService {
	return &productService{repo}
}
