package product

import (
	"time"

	"github.com/yescorihuela/agrak/domain/entity"
	"github.com/yescorihuela/agrak/domain/factory"
	"github.com/yescorihuela/agrak/domain/repository"
	"github.com/yescorihuela/agrak/infrastructure/postgresql"
	"github.com/yescorihuela/agrak/infrastructure/product/model"
)

type PersistenceProductRepository struct {
	Connection postgresql.PostgresqlRepository
}

func NewPersistenceProductRepository(conn postgresql.PostgresqlRepository) repository.ProductRepository {
	return &PersistenceProductRepository{
		Connection: conn,
	}
}

func (p *PersistenceProductRepository) Save(product entity.Product) error {
	db, err := p.Connection.GetConnection()
	if err != nil {
		return err
	}

	isValid, err := product.IsValid()

	if isValid {
		err := db.Create(model.ProductModel{
			Sku:            product.Sku,
			Name:           product.Name,
			Brand:          product.Brand,
			Size:           product.Size,
			Price:          product.Price,
			PrincipalImage: product.PrincipalImage.Url,
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
	result := db.First(&product, sku)
	if result.Error != nil {
		return nil, result.Error
	}
	entityProduct := factory.NewProduct(
		product.Sku,
		product.Name,
		product.Brand,
		product.Size,
		product.Price,
		entity.URLImage{Url: product.PrincipalImage},
		// []entity.URLImage{},
	)

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
		entityProducts = append(entityProducts,
			*factory.NewProduct(
				v.Sku,
				v.Name,
				v.Brand,
				v.Size,
				v.Price,
				entity.URLImage{Url: v.PrincipalImage},
				// []entity.URLImage{},
			),
		)
	}
	return entityProducts, nil
}

func (p *PersistenceProductRepository) Update(product entity.Product) error {
	return nil
}

func (p *PersistenceProductRepository) Delete(sku string) error {
	db, err := p.Connection.GetConnection()
	if err != nil {
		return err
	}
	result := db.Delete(&model.ProductModel{}, sku)
	if result.Error != nil {
		return err
	}
	return nil
}
