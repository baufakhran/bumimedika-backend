package customer

import (
	"bumimedika-backend/internal/domain/dto"

	"gorm.io/gorm"
)

func (s *custumerService) Manage(payload *dto.Customer) error {
	product, err := s.GetByID(payload.ID)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}
	if product == nil {
		productEntity := payload.MapProductDtoToEntity()
		err = s.repo.Create(productEntity)
		return err
	} else {
		productEntity := payload.MapProductDtoToEntity()
		productEntity.ID = product.ID
		return s.repo.Update(productEntity)
	}
}

func (s *custumerService) GetByID(id uint64) (*dto.Customer, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dto.ToCustomerDto(product), nil
}

func (s *custumerService) GetList(name, sortField, sortOrder string, page, limit int) ([]dto.Customer, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	listProducts, total, err := s.repo.GetList(name, sortField, sortOrder, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return dto.ToCustomerDtoList(listProducts), total, nil
}

func (s *custumerService) Delete(id uint64) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
