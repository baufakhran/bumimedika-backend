package customer

import (
	entities "bumimedika-backend/internal/domain/entities"

	"gorm.io/gorm"
)

type ICustumerRepository interface {
	// customer
	Create(product *entities.Customer) error
	Update(product *entities.Customer) error
	Delete(id uint64) error
	GetByID(id uint64) (*entities.Customer, error)
	GetList(name string, sortField, sortOrder string, limit, offset int) ([]entities.Customer, int64, error)

	// customer type
	CreateType(customerType *entities.CustomerType) error
	UpdateType(customerType *entities.CustomerType) error
	GetTypeByID(id uint64) (*entities.CustomerType, error)
	GetTypeList(typeName, sortField, sortOrder string, limit, offset int) ([]entities.CustomerType, int64, error)
	DeleteType(id uint64) error
}

type customerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository constructor
func InitCustomerRepository(db *gorm.DB) *customerRepository {
	return &customerRepository{db: db}
}
