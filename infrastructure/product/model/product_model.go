package model

import "time"

type ProductModel struct {
	Sku            string    `gorm:"column:sku;primaryKey"`
	Name           string    `gorm:"column:name;not null"`
	Brand          string    `gorm:"column:brand;not null"`
	Size           string    `gorm:"column:size;default:ST"`
	Price          float64   `gorm:"column:price"`
	PrincipalImage string    `gorm:"column:principal_image"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	OtherImages    []string  `gorm:"column:other_images"`
}

func (p *ProductModel) TableName() string {
	return "products"
}
