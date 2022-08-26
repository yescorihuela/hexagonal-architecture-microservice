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
	PrincipalImage string
	OtherImages    []string
}

func (p *Product) IsValid() (bool, error) {
	if strings.TrimSpace(p.Sku) == "" {
		return false, errors.New("empty sku")
	}
	if !p.IsSKUValid() {
		return false, errors.New("invalid sku format (right format: FAL-XXXXXXX)")
	}
	if strings.TrimSpace(p.Name) == "" {
		return false, errors.New("empty name")
	}
	if strings.TrimSpace(p.Brand) == "" {
		return false, errors.New("empty brand")
	}

	if strings.TrimSpace(p.PrincipalImage) == "" {
		return false, errors.New("principal image url empty")
	}
	if p.Price == 0.0 {
		return false, errors.New("price with zero value")
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
	correlative, err := strconv.Atoi(splittedSku[1])
	if err != nil {
		return false
	}

	if correlative < SkuMin || correlative > SkuMax {
		return false
	}
	return true
}
