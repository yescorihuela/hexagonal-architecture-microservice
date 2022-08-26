package product

import (
	"github.com/stretchr/testify/mock"
	"github.com/yescorihuela/agrak/domain/entity"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Save(product entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *RepositoryMock) GetBySku(sku string) (*entity.Product, error) {
	args := m.Called(sku)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *RepositoryMock) GetAllProducts() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (m *RepositoryMock) Update(oldSku string, product entity.Product) (*entity.Product, error) {
	args := m.Called(oldSku, product)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *RepositoryMock) Delete(sku string) error {
	args := m.Called(sku)
	return args.Error(0)
}
