package product

import (
	"errors"
	"strings"
	"time"

	"github.com/yescorihuela/agrak/domain/entity"
	"github.com/yescorihuela/agrak/domain/factory"
	"github.com/yescorihuela/agrak/domain/repository"
	"github.com/yescorihuela/agrak/infrastructure/database"
	"github.com/yescorihuela/agrak/infrastructure/postgresql/product/model"
	"github.com/yescorihuela/agrak/shared/common"
)

type PersistenceProductRepository struct {
	Connection database.GenericDatabaseRepository
}

func NewPersistenceProductRepository(conn database.GenericDatabaseRepository) repository.ProductRepository {
	return &PersistenceProductRepository{
		Connection: conn,
	}
}

func (p *PersistenceProductRepository) Save(product entity.Product) error {
	db, err := p.Connection.GetConnection()
	if err != nil {
		return err
	}

	checkedProduct, _ := p.GetBySku(product.Sku)
	if checkedProduct != nil {
		return errors.New("duplicated sku")
	}

	isValid, err := product.IsValid()

	if isValid {
		err := db.Create(model.ProductModel{
			Sku:            product.Sku,
			Name:           product.Name,
			Brand:          product.Brand,
			Size:           product.Size,
			Price:          product.Price,
			PrincipalImage: product.PrincipalImage,
			OtherImages:    common.GetStringFromSlicedUrls(product.OtherImages),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})

		if err != nil {
			return err.Error
		}
	}
	return err
}

func (p *PersistenceProductRepository) GetBySku(sku string) (*entity.Product, error) {
	db, err := p.Connection.GetConnection()
	if err != nil {
		return nil, err
	}
	product := model.ProductModel{}
	result := db.First(&product, "sku = ?", sku)
	if result.Error != nil {
		return nil, result.Error
	}
	otherImages := common.GetSlicedUrls(product.OtherImages)
	entityProduct, err := factory.NewProduct(
		product.Sku,
		product.Name,
		product.Brand,
		product.Size,
		product.Price,
		product.PrincipalImage,
		otherImages,
	)
	if err != nil {
		return nil, err
	}

	return entityProduct, nil
}

func (p *PersistenceProductRepository) GetAllProducts() ([]entity.Product, error) {
	db, err := p.Connection.GetConnection()
	if err != nil {
		return nil, err
	}

	products := make([]model.ProductModel, 0)

	result := db.Find(&products)
	if result.Error != nil {
		return nil, err
	}
	entityProducts := make([]entity.Product, 0)
	for _, v := range products {
		productFromModel := entity.Product{
			Sku:            v.Sku,
			Name:           v.Name,
			Brand:          v.Brand,
			Size:           v.Size,
			Price:          v.Price,
			PrincipalImage: v.PrincipalImage,
			OtherImages:    common.GetSlicedUrls(v.OtherImages),
		}
		entityProducts = append(entityProducts, productFromModel)
	}
	return entityProducts, nil
}

func (p *PersistenceProductRepository) Update(oldSku string, product entity.Product) (*entity.Product, error) {
	db, err := p.Connection.GetConnection()
	if err != nil {
		return nil, err
	}
	oldProduct := model.ProductModel{
		Sku: oldSku,
	}

	otherImages := strings.Join(product.OtherImages, ",")
	newProduct := model.ProductModel{
		Sku:            product.Sku,
		Name:           product.Name,
		Brand:          product.Brand,
		Size:           product.Size,
		Price:          product.Price,
		PrincipalImage: product.PrincipalImage,
		OtherImages:    otherImages,
		UpdatedAt:      time.Now(),
	}

	result := db.Model(&oldProduct).Updates(newProduct)
	if result.Error != nil {
		return nil, result.Error
	}

	updatedProduct, err := factory.NewProduct(
		newProduct.Sku,
		newProduct.Name,
		newProduct.Brand,
		newProduct.Size,
		newProduct.Price,
		newProduct.PrincipalImage,
		common.GetSlicedUrls(otherImages),
	)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (p *PersistenceProductRepository) Delete(sku string) error {
	db, err := p.Connection.GetConnection()
	if err != nil {
		return err
	}
	result := db.Delete(&model.ProductModel{}, "sku = ?", sku)
	if result.Error != nil {
		return err
	}
	return nil
}
