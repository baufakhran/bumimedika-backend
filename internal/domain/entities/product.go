package entities

import "time"

type Product struct {
	ID            uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string     `json:"name" gorm:"size:255;not null"`
	BatchNumber   string     `json:"batch_number" gorm:"size:100;not null"`
	PurchasePrice float64    `json:"purchase_price" gorm:"type:decimal(15,2);not null"`
	Price         float64    `json:"price" gorm:"type:decimal(15,2);not null"`
	Stock         uint       `json:"stock" gorm:"default:0"`
	Sold          uint       `json:"sold" gorm:"default:0"`
	Unit          string     `json:"unit" gorm:"size:50;not null;default:''"`
	Category      string     `json:"category" gorm:"size:100;not null;default:''"`
	Location      string     `json:"location" gorm:"size:100;not null;default:''"`
	Color         string     `json:"color" gorm:"size:50;not null;default:''"`
	Size          string     `json:"size" gorm:"size:50;not null;default:''"`
	Brand         string     `json:"brand" gorm:"size:100;not null;default:''"`
	ExpiredAt     *time.Time `json:"expired_at" gorm:"type:date"`
	CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
