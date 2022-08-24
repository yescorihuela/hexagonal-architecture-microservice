package repository

import "github.com/yescorihuela/agrak/domain/entity"

type ProductRepository interface {
	Save() (*entity.Product, error)
	Update() error
	GetBySku(sku string) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	Delete(sku string) error
}
