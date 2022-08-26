package application

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/agrak/domain/entity"
	"github.com/yescorihuela/agrak/domain/factory"
	"github.com/yescorihuela/agrak/infrastructure/response"
	"github.com/yescorihuela/agrak/usecase"
)

type ProductHandlers struct {
	service usecase.Service
}

func NewProductHandlers(service usecase.Service) *ProductHandlers {
	return &ProductHandlers{
		service: service,
	}
}

func (ph *ProductHandlers) GetProductBySku(ctx *gin.Context) {
	sku := ctx.Param("sku")
	product, err := ph.service.FindBySku(sku)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	response := response.ConvertFromEntityToResponse(*product)
	ctx.JSON(http.StatusOK, response)
}

func (ph *ProductHandlers) CreateProduct(ctx *gin.Context) {
	reqProduct := struct {
		Sku            string  `json:"sku"`
		Name           string  `json:"name"`
		Brand          string  `json:"brand"`
		Size           string  `json:"size"`
		Price          float64 `json:"price"`
		PrincipalImage string  `json:"principal_image"`
	}{}
	err := ctx.ShouldBindJSON(&reqProduct)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, nil)
	}
	product := factory.NewProduct(
		reqProduct.Sku,
		reqProduct.Name,
		reqProduct.Brand,
		reqProduct.Size,
		reqProduct.Price,
		entity.URLImage{Url: reqProduct.PrincipalImage},
	)
	if product != nil {
		if validProduct, _ := product.IsValid(); validProduct {
			err = ph.service.CreateProduct(*product)
			if err != nil {
				ctx.JSON(http.StatusUnprocessableEntity, err)
				return
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, err)
			return
		}
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	response := response.ConvertFromEntityToResponse(*product)
	ctx.JSON(http.StatusCreated, response)
}

func (ph *ProductHandlers) GetAllProducts(ctx *gin.Context) {
	products, err := ph.service.FindAll()
	responseJSON := make([]response.ProductResponse, 0)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	for _, product := range products {
		responseJSON = append(responseJSON, *response.ConvertFromEntityToResponse(product))
	}
	ctx.JSON(http.StatusOK, responseJSON)
}

func (ph *ProductHandlers) UpdateProduct(ctx *gin.Context) {
	sku := ctx.Param("sku")
	product, err := ph.service.FindBySku(sku)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	reqProduct := struct {
		Sku            string  `json:"sku"`
		Name           string  `json:"name"`
		Brand          string  `json:"brand"`
		Size           string  `json:"size"`
		Price          float64 `json:"price"`
		PrincipalImage string  `json:"principal_image"`
	}{}
	err = ctx.ShouldBindJSON(&reqProduct)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, nil)
	}
	newProduct := factory.NewProduct(
		reqProduct.Sku,
		reqProduct.Name,
		reqProduct.Brand,
		reqProduct.Size,
		reqProduct.Price,
		entity.URLImage{Url: reqProduct.PrincipalImage},
	)
	if !reflect.DeepEqual(product, newProduct) {
		product, err = ph.service.UpdateProduct(sku, *newProduct)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, nil)
			return
		}

		if value, _ := product.IsValid(); value {
			ctx.JSON(http.StatusOK, response.ConvertFromEntityToResponse(*product))
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, response.ConvertFromEntityToResponse(*product))
			return
		}
	}
	ctx.JSON(http.StatusOK, response.ConvertFromEntityToResponse(*newProduct))
}

func (ph *ProductHandlers) Delete(ctx *gin.Context) {
	sku := ctx.Param("sku")
	err := ph.service.DeleteProduct(sku)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
