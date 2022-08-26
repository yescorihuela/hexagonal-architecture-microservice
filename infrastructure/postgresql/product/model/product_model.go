package model

import "time"

type ProductModel struct {
	Sku            string    `gorm:"column:sku;primaryKey"`
	Name           string    `gorm:"column:name;not null"`
	Brand          string    `gorm:"column:brand;not null"`
	Size           string    `gorm:"column:size;default:ST"`
	Price          float64   `gorm:"column:price;scale:10,precision:2"`
	PrincipalImage string    `gorm:"column:principal_image"`
	OtherImages    string    `gorm:"column:other_images"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (p *ProductModel) TableName() string {
	return "products"
}
