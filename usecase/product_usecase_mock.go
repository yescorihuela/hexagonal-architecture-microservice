package usecase

import (
	"github.com/stretchr/testify/mock"
	"github.com/yescorihuela/agrak/domain/entity"
)

type UseCaseMock struct {
	mock.Mock
}

func (m *UseCaseMock) CreateProduct(product entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *UseCaseMock) FindBySku(sku string) (*entity.Product, error) {
	args := m.Called(sku)
	var mockedEntityProduct *entity.Product
	var mockedError error
	if args.Get(0) != nil {
		mockedEntityProduct = args.Get(0).(*entity.Product)
	}

	if args.Get(1) != nil {
		mockedError = args.Get(1).(error)
	}

	return mockedEntityProduct, mockedError
}

func (m *UseCaseMock) FindAll() ([]entity.Product, error) {
	args := m.Called()
	var mockedEntityProduct []entity.Product
	var mockedError error
	if args.Get(0) != nil {
		mockedEntityProduct = args.Get(0).([]entity.Product)
	}

	if args.Get(1) != nil {
		mockedError = args.Get(1).(error)
	}

	return mockedEntityProduct, mockedError
}

func (m *UseCaseMock) UpdateProduct(oldSku string, product entity.Product) (*entity.Product, error) {
	args := m.Called(oldSku, product)
	var mockedEntityProduct *entity.Product
	var mockedError error
	if args.Get(0) != nil {
		mockedEntityProduct = args.Get(0).(*entity.Product)
	}

	if args.Get(1) != nil {
		mockedError = args.Get(1).(error)
	}

	return mockedEntityProduct, mockedError
}

func (m *UseCaseMock) DeleteProduct(sku string) error {
	args := m.Called(sku)
	return args.Error(0)
}
