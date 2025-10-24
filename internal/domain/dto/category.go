package dto

import (
	entities "bumimedika-backend/internal/domain/entities"
	"time"
)

type CustomerType struct {
	ID        uint64     `json:"id"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy uint64     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *uint64    `json:"updated_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *uint64    `json:"deleted_by,omitempty"`
}

type Customer struct {
	ID             uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string     `json:"name"`
	CustomerTypeID uint64     `json:"customer_type_id"`
	Pharmacist     string     `json:"pharmacist,omitempty"`
	SipaNumber     string     `json:"sipa_number,omitempty"`
	SipaExpiredAt  *time.Time `json:"sipa_expired_at,omitempty"`
	Address        string     `json:"address,omitempty"`
	Phone          string     `json:"phone,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	CreatedBy      uint64     `json:"created_by"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	UpdatedBy      *uint64    `json:"updated_by,omitempty"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	DeletedBy      *uint64    `json:"deleted_by,omitempty"`
}

func (c *Customer) MapProductDtoToEntity() *entities.Customer {
	return &entities.Customer{
		ID:             c.ID,
		Name:           c.Name,
		CustomerTypeID: c.CustomerTypeID,
		Pharmacist:     c.Pharmacist,
		SipaNumber:     c.SipaNumber,
		SipaExpiredAt:  c.SipaExpiredAt,
		Address:        c.Address,
		Phone:          c.Phone,
		CreatedAt:      c.CreatedAt,
		CreatedBy:      c.CreatedBy,
		UpdatedAt:      c.UpdatedAt,
		UpdatedBy:      c.UpdatedBy,
		DeletedAt:      c.DeletedAt,
		DeletedBy:      c.DeletedBy,
	}
}

func ToCustomerDto(entity *entities.Customer) *Customer {
	if entity == nil {
		return nil
	}
	return &Customer{
		ID:             entity.ID,
		Name:           entity.Name,
		CustomerTypeID: entity.CustomerTypeID,
		Pharmacist:     entity.Pharmacist,
		SipaNumber:     entity.SipaNumber,
		SipaExpiredAt:  entity.SipaExpiredAt,
		Address:        entity.Address,
		Phone:          entity.Phone,
		CreatedAt:      entity.CreatedAt,
		CreatedBy:      entity.CreatedBy,
		UpdatedAt:      entity.UpdatedAt,
		UpdatedBy:      entity.UpdatedBy,
		DeletedAt:      entity.DeletedAt,
		DeletedBy:      entity.DeletedBy,
	}
}

func ToCustomerDtoList(entities []entities.Customer) []Customer {
	var dtos []Customer
	for _, entity := range entities {
		dtos = append(dtos, *ToCustomerDto(&entity))
	}
	return dtos
}
