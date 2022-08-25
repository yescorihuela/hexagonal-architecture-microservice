package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/agrak/infrastructure/response"
	"github.com/yescorihuela/agrak/usecase"
)

type ProductHandlers struct {
	service usecase.ProductService
}

func NewProductHandlers(service usecase.ProductService) *ProductHandlers {
	return &ProductHandlers{
		service: service,
	}
}

func (ph *ProductHandlers) GetProductBySku(ctx *gin.Context) {
	product, err := ph.service.FindBySku("")
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	response := response.ConvertFromEntityToResponse(*product)
	ctx.JSON(http.StatusOK, response)
}
