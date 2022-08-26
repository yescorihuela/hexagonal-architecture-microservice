package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yescorihuela/agrak/domain/entity"
	"github.com/yescorihuela/agrak/infrastructure/postgresql/product"
)

func TestProductService_Save(t *testing.T) {
	t.Run("should not return an error", func(t *testing.T) {
		productRepositoryMock := new(product.RepositoryMock)
		productFake := entity.Product{
			Sku:            "FAL-9999999",
			Name:           "Bicicleta infantil",
			Brand:          "Oxford",
			Size:           "16",
			Price:          130000.00,
			PrincipalImage: "https://via.placeholder.com/500x500.png?text=Principal+image",
			OtherImages: []string{
				"https://via.placeholder.com/728x190.png?text=Agrak+Exercise+Resolution",
				"https://via.placeholder.com/500x260.png?text=Agrak+Exercise+Resolution",
				"https://via.placeholder.com/500x500.png?text=Agrak+Exercise+Resolution",
			},
		}
		productRepositoryMock.On("Save", productFake).Return(nil)

		useCase := NewProductService(productRepositoryMock)
		err := useCase.CreateProduct(productFake)
		assert.NoError(t, err)
	})
	t.Run("should return an error", func(t *testing.T) {
		t.Run("should not return an error", func(t *testing.T) {
			productRepositoryMock := new(product.RepositoryMock)
			productRepositoryMock.On("Save", mock.Anything).Return(errors.New("any repository error"))

			useCase := NewProductService(productRepositoryMock)
			err := useCase.CreateProduct(entity.Product{})
			assert.EqualError(t, err, "any repository error")
		})
	})
}
