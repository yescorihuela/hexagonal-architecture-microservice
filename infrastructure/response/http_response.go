package response

import (
	"github.com/yescorihuela/agrak/domain/entity"
)

type DTOProduct struct {
	Sku            string   `json:"sku"`
	Name           string   `json:"name"`
	Brand          string   `json:"brand"`
	Size           string   `json:"size"`
	Price          float64  `json:"price"`
	PrincipalImage string   `json:"principal_image"`
	OtherImages    []string `json:"other_images"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}

func ConvertFromEntityToResponse(ep entity.Product) *DTOProduct {
	return &DTOProduct{
		Sku:            ep.Sku,
		Name:           ep.Name,
		Brand:          ep.Brand,
		Size:           ep.Size,
		Price:          ep.Price,
		PrincipalImage: ep.PrincipalImage,
		OtherImages:    ep.OtherImages,
	}
}
