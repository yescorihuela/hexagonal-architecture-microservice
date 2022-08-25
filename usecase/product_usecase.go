package usecase

import (
	"github.com/yescorihuela/agrak/domain/entity"
	"github.com/yescorihuela/agrak/domain/repository"
)

type Service interface {
	Create() error
	FindBySku(sku string) (*entity.Product, error)
	FindAll() error
	Update() error
	Delete(sku string) error
}

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) Service {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) Create() error {
	return nil
}

func (s *ProductService) FindBySku(sku string) (*entity.Product, error) {
	product, err := s.repository.GetBySku(sku)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) FindAll() error {
	s.repository.GetAllProducts()
	return nil
}

func (s *ProductService) Update() error {
	return nil
}

func (s *ProductService) Delete(sku string) error {
	s.repository.Delete(sku)
	return nil
}
