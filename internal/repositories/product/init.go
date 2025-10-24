package product

import (
	entities "bumimedika-backend/internal/domain/entities"

	"gorm.io/gorm"
)

type IProductRepository interface {
	Create(product *entities.Product) error
	Update(product *entities.Product) error
	Delete(id uint64) error
	GetByID(id uint64) (*entities.Product, error)
	GetList(name string, sortField, sortOrder string, limit, offset int) ([]entities.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func InitProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}
