package application

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/agrak/domain/factory"
	"github.com/yescorihuela/agrak/infrastructure/response"
	"github.com/yescorihuela/agrak/usecase"
)

var request = struct {
	Sku            string   `json:"sku"`
	Name           string   `json:"name"`
	Brand          string   `json:"brand"`
	Size           string   `json:"size"`
	Price          float64  `json:"price"`
	PrincipalImage string   `json:"principal_image"`
	OtherImages    []string `json:"other_images"`
}{}

type ProductHandlers struct {
	service usecase.Service
}

func NewProductHandlers(service usecase.Service) *ProductHandlers {
	return &ProductHandlers{
		service: service,
	}
}

// GetProductBySku godoc
// @Summary Retrieve a product by SKU
// @Description get product by SKU as json
// @Accept json
// @Produce json
// @param sku path string true "Product unique SKU"
// @Success 200 {object} response.DTOProduct
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/products/{sku} [get]
func (ph *ProductHandlers) GetProductBySku(ctx *gin.Context) {
	sku := ctx.Param("sku")
	product, err := ph.service.FindBySku(sku)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.NewErrorResponse(err.Error()))
		return
	}

	response := response.ConvertFromEntityToResponse(*product)
	ctx.JSON(http.StatusOK, response)
}

// CreateProduct godoc
// @Summary Add a product
// @Description add by json product
// @Accept json
// @Produce json
// @param product body response.DTOProduct true "Add new product with unique SKU"
// @Success 201 {object} response.DTOProduct
// @Failure 422 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/products/ [post]
func (ph *ProductHandlers) CreateProduct(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
	}
	product, err := factory.NewProduct(
		request.Sku,
		request.Name,
		request.Brand,
		request.Size,
		request.Price,
		request.PrincipalImage,
		request.OtherImages,
	)

	if product != nil {
		if validProduct, err := product.IsValid(); validProduct {
			err = ph.service.CreateProduct(*product)
			if err != nil {
				ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
				return
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
			return
		}
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
		return
	}
	response := response.ConvertFromEntityToResponse(*product)
	ctx.JSON(http.StatusCreated, response)
}

// GetAllProducts godoc
// @Summary List all the stored products
// @Description list all the products as an array
// @Accept json
// @Produce json
// @param none query string false "Not required."
// @Success 200 {array} response.DTOProduct
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/products/ [get]
func (ph *ProductHandlers) GetAllProducts(ctx *gin.Context) {
	products, err := ph.service.FindAll()
	responseJSON := make([]response.DTOProduct, 0)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.NewErrorResponse(err.Error()))
		return
	}
	for _, product := range products {
		responseJSON = append(responseJSON, *response.ConvertFromEntityToResponse(product))
	}
	ctx.JSON(http.StatusOK, responseJSON)
}

// UpdateProduct godoc
// @Summary Delete a product by SKU
// @Description delete product by SKU
// @Accept json
// @Produce json
// @param sku path string true "Product unique SKU"
// @param product body response.DTOProduct true "Product body with unique SKU"
// @Success 200 {object} response.DTOProduct
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/products/{sku} [put]
func (ph *ProductHandlers) UpdateProduct(ctx *gin.Context) {
	sku := ctx.Param("sku")
	product, err := ph.service.FindBySku(sku)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.NewErrorResponse(err.Error()))
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
		return
	}

	newProduct, err := factory.NewProduct(
		request.Sku,
		request.Name,
		request.Brand,
		request.Size,
		request.Price,
		request.PrincipalImage,
		request.OtherImages,
	)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
		return
	}

	if !reflect.DeepEqual(product, newProduct) {
		product, err = ph.service.UpdateProduct(sku, *newProduct)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
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

// Delete godoc
// @Summary Delete a product by SKU
// @Description delete product by SKU
// @Accept json
// @Produce json
// @param sku path string true "Product unique SKU"
// @Success 204 {object} nil
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/products/{sku} [delete]
func (ph *ProductHandlers) Delete(ctx *gin.Context) {
	sku := ctx.Param("sku")
	err := ph.service.DeleteProduct(sku)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
