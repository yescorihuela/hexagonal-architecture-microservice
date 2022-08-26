package usecase

import (
	"github.com/yescorihuela/agrak/domain/entity"
	"github.com/yescorihuela/agrak/domain/repository"
)

type Service interface {
	CreateProduct(product entity.Product) error
	FindBySku(sku string) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
	UpdateProduct(oldSku string, product entity.Product) (*entity.Product, error)
	DeleteProduct(sku string) error
}

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) Service {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) CreateProduct(product entity.Product) error {
	err := s.repository.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) FindBySku(sku string) (*entity.Product, error) {
	product, err := s.repository.GetBySku(sku)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) FindAll() ([]entity.Product, error) {
	products, err := s.repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) UpdateProduct(oldSku string, product entity.Product) (*entity.Product, error) {
	updatedProduct, err := s.repository.Update(oldSku, product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (s *ProductService) DeleteProduct(sku string) error {
	err := s.repository.Delete(sku)
	if err != nil {
		return err
	}
	return nil
}
