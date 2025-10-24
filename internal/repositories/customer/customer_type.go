package customer

import (
	entities "bumimedika-backend/internal/domain/entities"
	"fmt"

	"gorm.io/gorm"
)

// ---------- CustomerType ----------

func (r *customerRepository) CreateType(customerType *entities.CustomerType) error {
	return r.db.Create(customerType).Error
}

func (r *customerRepository) UpdateType(customerType *entities.CustomerType) error {
	return r.db.Updates(customerType).Error
}

func (r *customerRepository) GetTypeByID(id uint64) (*entities.CustomerType, error) {
	var customerType entities.CustomerType
	if err := r.db.First(&customerType, id).Error; err != nil {
		return nil, err
	}
	return &customerType, nil
}

// GetTypeList supports filtering by type name, sorting, and pagination
func (r *customerRepository) GetTypeList(typeName, sortField, sortOrder string, limit, offset int) ([]entities.CustomerType, int64, error) {
	var types []entities.CustomerType
	var total int64

	query := r.db.Model(&entities.CustomerType{})

	if typeName != "" {
		query = query.Where("type LIKE ?", "%"+typeName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSortFields := map[string]bool{
		"id": true, "type": true, "created_at": true,
		"updated_at": true, "deleted_at": true,
	}

	if !allowedSortFields[sortField] {
		sortField = "id"
	}

	order := "ASC"
	if sortOrder == "desc" {
		order = "DESC"
	}

	query = query.Order(fmt.Sprintf("%s %s", sortField, order))

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Find(&types).Error; err != nil {
		return nil, 0, err
	}

	return types, total, nil
}

func (r *customerRepository) DeleteType(id uint64) error {
	return r.db.Model(&entities.CustomerType{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}
