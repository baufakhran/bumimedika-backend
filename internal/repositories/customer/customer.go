package customer

import (
	entities "bumimedika-backend/internal/domain/entities"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ---------- Customer ----------

func (r *customerRepository) Create(customer *entities.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) Update(customer *entities.Customer) error {
	return r.db.Updates(customer).Error
}

func (r *customerRepository) GetByID(id uint64) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetList supports filtering by name, sorting, and pagination
func (r *customerRepository) GetList(name, sortField, sortOrder string, limit, offset int) ([]entities.Customer, int64, error) {
	var customers []entities.Customer
	var total int64

	query := r.db.Model(&entities.Customer{}).Where("status = ?", "1")

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSortFields := map[string]bool{
		"id": true, "name": true, "created_at": true,
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

	if err := query.Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (r *customerRepository) Delete(id uint64) error {
	return r.db.Model(&entities.Customer{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_at": gorm.DeletedAt{Time: time.Now(), Valid: true},
		"status":     0,
	}).Error
}
