package customer

import (
	"bumimedika-backend/internal/domain/dto"
	repository "bumimedika-backend/internal/repositories/customer"
)

type ICustumerService interface {
	Manage(product *dto.Customer) error
	GetByID(id uint64) (*dto.Customer, error)
	Delete(id uint64) error
	GetList(name, sortField, sortOrder string, page, limit int) ([]dto.Customer, int64, error)
}

type custumerService struct {
	repo repository.ICustumerRepository
}

func InitCustumerService(repo repository.ICustumerRepository) ICustumerService {
	return &custumerService{repo}
}
