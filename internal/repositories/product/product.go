package product

import (
	entities "bumimedika-backend/internal/domain/entities"
	"fmt"
)

func (r *productRepository) Create(product *entities.Product) error {
	err := r.db.Create(product).Error
	return err
}

func (r *productRepository) Update(product *entities.Product) error {
	return r.db.Updates(product).Error

}

func (r *productRepository) GetByID(id uint64) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetList supports filtering by name, sorting, and pagination.
func (r *productRepository) GetList(name string, sortField, sortOrder string, limit, offset int) ([]entities.Product, int64, error) {
	var products []entities.Product
	var total int64

	query := r.db.Model(&entities.Product{})

	// Filtering
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// Count total before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sort whitelist — protect from SQL injection
	allowedSortFields := map[string]bool{
		"id": true, "name": true, "purchase_price": true,
		"selling_price": true, "stock": true, "sold": true,
		"expired_at": true, "brand": true,
	}

	if !allowedSortFields[sortField] {
		sortField = "id"
	}

	order := "ASC"
	if sortOrder == "desc" {
		order = "DESC"
	}

	query = query.Order(fmt.Sprintf("%s %s", sortField, order))

	// Apply pagination
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
