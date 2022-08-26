package factory

import (
	"errors"
	"fmt"
	"strings"

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
) (*entity.Product, error) {
	if validateLengthString(sku, 3, 50) {
		return nil, ErrorFieldLimit("sku", 3, 50)
	}

	if validateLengthString(name, 3, 50) {
		return nil, ErrorFieldLimit("name", 3, 50)
	}

	if validateLengthString(brand, 3, 50) {
		return nil, ErrorFieldLimit("brand", 3, 50)
	}

	if strings.TrimSpace(size) != "" {
		if validateLengthString(size, 1, 15) {
			return nil, ErrorFieldLimit("size", 3, 50)
		}
	} else {
		size = "ST"
	}

	if price < entity.PriceMin || price > entity.PriceMax {
		return nil, ErrorFieldLimit("price", entity.PriceMin, entity.PriceMax)
	}

	if strings.TrimSpace(principalImage.Url) == "" {
		return nil, errors.New("principal_url must be a valid url")
	}
	return &entity.Product{
		Sku:            sku,
		Name:           name,
		Brand:          brand,
		Size:           size,
		Price:          price,
		PrincipalImage: principalImage,
		// OtherImage:     otherImages,
	}, nil
}

func validateLengthString(value string, minLength, maxLength int) bool {
	if len(value) < minLength || len(value) > maxLength {
		return true
	}
	return false
}

func ErrorFieldLimit(fieldName string, minLength, maxLength int) error {
	return fmt.Errorf("%s must be between %d and %d", fieldName, minLength, maxLength)
}
