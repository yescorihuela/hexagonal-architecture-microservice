package factory

import (
	"fmt"

	"github.com/yescorihuela/agrak/domain/entity"
)

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
		return nil
	}

	if validateString(name, 3, 50) {
		return nil
	}

	if validateString(brand, 3, 50) {
		return nil
	}

	if validateString(size, 1, 15) {
		return nil
	}

	if price < entity.PriceMin || price > entity.PriceMax {
		return nil
	}
	fmt.Println("price", price)
	if principalImage.Url == "" {
		return nil
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
