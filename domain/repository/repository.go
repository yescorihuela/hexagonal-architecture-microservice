package repository

import "github.com/yescorihuela/agrak/domain/entity"

type ProductRepository interface {
	Save(p entity.Product) error
	Update(p entity.Product) error
	GetBySku(sku string) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	Delete(sku string) error
}
