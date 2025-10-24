package entities

import "time"

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
