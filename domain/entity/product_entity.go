package entity

import (
	"strconv"
	"strings"
)

const (
	SKU_MIN    = 1_000_000
	SKU_MAX    = 9_999_999
	SKU_PREFIX = "FAL"
)

type Product struct {
	Sku            string
	Name           string
	Brand          string
	Size           string
	Price          float64
	PrincipalImage URLImage
	OtherImage     []URLImage
}

func (p *Product) IsValid() bool {
	if strings.TrimSpace(p.Sku) == "" {
		return false
	}
	if strings.TrimSpace(p.Name) == "" {
		return false
	}
	if strings.TrimSpace(p.Brand) == "" {
		return false
	}
	if strings.TrimSpace(p.Size) == "" {
		return false
	}
	if strings.TrimSpace(string(p.PrincipalImage.Url)) == "" {
		return false
	}
	if p.Price == 0.0 {
		return false
	}
	for _, v := range p.OtherImage {
		if strings.TrimSpace(string(v.Url)) == "" {
			return false
		}
	}
	return true
}

func (p *Product) IsSKUValid() bool {
	splittedSku := strings.Split(p.Sku, "-")
	if splittedSku[0] != SKU_PREFIX {
		return false
	}
	correlative, err := strconv.Atoi(splittedSku[1])
	if err != nil {
		return false
	}

	if correlative < SKU_MIN || correlative > SKU_MAX {
		return false
	}
	return true
}
