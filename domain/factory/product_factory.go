package factory

import "github.com/yescorihuela/agrak/domain/entity"

func NewProduct(
	sku,
	name,
	brand,
	size string,
	price float64,
	principalImage entity.URLImage,
	// otherImages []entity.URLImage,
) *entity.Product {
	if validateString(sku, 3, 50) {
		return &entity.Product{}
	}
	if validateString(name, 3, 50) {
		return &entity.Product{}
	}
	if validateString(brand, 3, 50) {
		return &entity.Product{}
	}
	if validateString(size, 3, 50) {
		return &entity.Product{}
	}
	if price < entity.PRICE_MIN || price > entity.PRICE_MAX {
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
		// OtherImage:     otherImages,
	}
}

func validateString(value string, minLength, maxLength int) bool {
	if len(value) < minLength || len(value) > maxLength {
		return true
	}
	return false
}
