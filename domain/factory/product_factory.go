package factory

import "github.com/yescorihuela/agrak/domain/entity"

func NewProduct(
	sku,
	name,
	brand,
	size string,
	price float64,
	principalImage entity.URLImage,
	otherImages []entity.URLImage,
) *entity.Product {
	if sku == "" {
		return &entity.Product{}
	}
	if name == "" {
		return &entity.Product{}
	}
	if brand == "" {
		return &entity.Product{}
	}
	if size == "" {
		return &entity.Product{}
	}
	if price <= 0.0 {
		return &entity.Product{}
	}
	if principalImage.Url == "" {
		return &entity.Product{}
	}
	return &entity.Product{
		Sku:            sku,
		Name:           name,
		Brand:          brand,
		Size:           size,
		Price:          price,
		PrincipalImage: principalImage,
		OtherImage:     otherImages,
	}
}
