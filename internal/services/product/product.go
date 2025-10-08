package product

import (
	"bumimedika-backend/internal/domain/dto"

	"gorm.io/gorm"
)

func (s *productService) Manage(payload *dto.Product) error {
	product, err := s.GetByID(payload.ID)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}
	if product == nil {
		productEntity := payload.MapProductDtoToEntity()
		return s.repo.Create(productEntity)
	} else {
		productEntity := payload.MapProductDtoToEntity()
		productEntity.ID = product.ID
		return s.repo.Update(productEntity)
	}
}

func (s *productService) GetByID(id uint64) (*dto.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dto.ToProductDto(product), nil
}

func (s *productService) GetList(name, sortField, sortOrder string, page, limit int) ([]dto.Product, int64, error) {
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

	return dto.ToProductDtoList(listProducts), total, nil
}
