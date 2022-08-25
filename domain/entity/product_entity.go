package entity

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	PriceMin  = 1.00
	PriceMax  = 99_999_999.00
	SkuMin    = 1_000_000
	SkuMax    = 9_999_999
	SkuPrefix = "FAL"
)

type Product struct {
	Sku            string
	Name           string
	Brand          string
	Size           string
	Price          float64
	PrincipalImage URLImage
	// OtherImage     []URLImage
}

func (p *Product) IsValid() (bool, error) {
	if strings.TrimSpace(p.Sku) == "" {
		return false, nil
	}
	if strings.TrimSpace(p.Name) == "" {
		return false, errors.New("")
	}
	if strings.TrimSpace(p.Brand) == "" {
		return false, errors.New("")
	}
	if strings.TrimSpace(p.Size) == "" {
		return false, errors.New("")
	}
	if strings.TrimSpace(string(p.PrincipalImage.Url)) == "" {
		return false, errors.New("")
	}
	if p.Price == 0.0 {
		return false, errors.New("")
	}
	// for _, v := range p.OtherImage {
	// 	if strings.TrimSpace(string(v.Url)) == "" {
	// 		return false, errors.New("")
	// 	}
	// }
	return true, nil
}

func (p *Product) IsSKUValid() bool {
	if matched, _ := regexp.MatchString(`FAL-[0-9]{7}`, p.Sku); !matched {
		return false
	}
	splittedSku := strings.Split(p.Sku, "-")
	if splittedSku[0] != SkuPrefix {
		return false
	}
	correlative, err := strconv.Atoi(splittedSku[1])
	if err != nil {
		return false
	}

	if correlative < SkuMin || correlative > SkuMax {
		return false
	}
	return true
}
