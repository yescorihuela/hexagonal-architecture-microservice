package response

import (
	"time"

	"github.com/yescorihuela/agrak/domain/entity"
)

type ProductResponse struct {
	Sku            string    `json:"sku"`
	Name           string    `json:"name"`
	Brand          string    `json:"brand"`
	Size           string    `json:"size"`
	Price          float64   `json:"price"`
	PrincipalImage string    `json:"principal_image"`
	OtherImages    []string  `json:"-"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func ConvertFromEntityToResponse(ep entity.Product) *ProductResponse {

	return &ProductResponse{
		Sku:            ep.Sku,
		Name:           ep.Name,
		Brand:          ep.Brand,
		Size:           ep.Size,
		Price:          ep.Price,
		PrincipalImage: ep.PrincipalImage.Url,
	}
}