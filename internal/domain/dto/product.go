package dto

import (
	entities "bumimedika-backend/internal/domain/entities"
	"time"
)

type Product struct {
	ID            uint64     `json:"id"`
	Name          string     `json:"name"`
	BatchNumber   string     `json:"batch_number"`
	PurchasePrice float64    `json:"purchase_price"`
	Price         float64    `json:"price"`
	Stock         uint       `json:"stock"`
	Sold          uint       `json:"sold"`
	Unit          string     `json:"unit"`
	Category      string     `json:"category"`
	Location      string     `json:"location"`
	Color         string     `json:"color"`
	Size          string     `json:"size"`
	Brand         string     `json:"brand"`
	ExpiredAt     *time.Time `json:"expired_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func ToProductDto(p *entities.Product) *Product {
	if p == nil {
		return nil
	}
	return &Product{
		ID:            p.ID,
		Name:          p.Name,
		BatchNumber:   p.BatchNumber,
		PurchasePrice: p.PurchasePrice,
		Price:         p.Price,
		Stock:         p.Stock,
		Sold:          p.Sold,
		Unit:          p.Unit,
		Category:      p.Category,
		Location:      p.Location,
		Color:         p.Color,
		Size:          p.Size,
		Brand:         p.Brand,
		ExpiredAt:     p.ExpiredAt,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func (p *Product) MapProductDtoToEntity() *entities.Product {
	if p == nil {
		return nil
	}
	return &entities.Product{
		ID:            p.ID,
		Name:          p.Name,
		BatchNumber:   p.BatchNumber,
		PurchasePrice: p.PurchasePrice,
		Price:         p.Price,
		Stock:         p.Stock,
		Sold:          p.Sold,
		Unit:          p.Unit,
		Category:      p.Category,
		Location:      p.Location,
		Color:         p.Color,
		Size:          p.Size,
		Brand:         p.Brand,
		ExpiredAt:     p.ExpiredAt,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func ToProductDtoList(products []entities.Product) []Product {
	var productDtos []Product
	for _, p := range products {
		productDtos = append(productDtos, *ToProductDto(&p))
	}
	return productDtos
}
