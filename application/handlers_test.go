package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yescorihuela/agrak/domain/entity"
	"github.com/yescorihuela/agrak/domain/factory"
	"github.com/yescorihuela/agrak/infrastructure/response"
	"github.com/yescorihuela/agrak/usecase"
)

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("CreateProduct - 201 Created", func(t *testing.T) {
		mockProductPayload := struct {
			Sku            string   `json:"sku"`
			Name           string   `json:"name"`
			Brand          string   `json:"brand"`
			Size           string   `json:"size"`
			Price          float64  `json:"price"`
			PrincipalImage string   `json:"principal_image"`
			OtherImages    []string `json:"other_images"`
		}{
			Sku:            "FAL-1000000",
			Name:           "Polera",
			Brand:          "CAT",
			Size:           "XL",
			Price:          20000.00,
			PrincipalImage: "https://placehold.jp/3d4070/ffffff/150x150.png",
			OtherImages: []string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		}
		mockEntityProduct, _ := factory.NewProduct(
			mockProductPayload.Sku,
			mockProductPayload.Name,
			mockProductPayload.Brand,
			mockProductPayload.Size,
			mockProductPayload.Price,
			mockProductPayload.PrincipalImage,
			mockProductPayload.OtherImages,
		)

		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("CreateProduct", *mockEntityProduct).Return(nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.POST("/products", NewProductHandlers(mockUsecase).CreateProduct)
		payload, _ := json.Marshal(mockProductPayload)
		request, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(payload))

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(mockProductPayload)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertExpectations(t)
	})

	t.Run("CreateProduct - 422 Unprocessable Entity", func(t *testing.T) {
		mockProductPayload := struct {
			Sku            string   `json:"sku"`
			Name           string   `json:"name"`
			Brand          string   `json:"brand"`
			Size           string   `json:"size"`
			Price          float64  `json:"price"`
			PrincipalImage string   `json:"principal_image"`
			OtherImages    []string `json:"other_images"`
		}{
			Sku:            "FAL-100000",
			Name:           "Polera",
			Brand:          "CAT",
			Size:           "XL",
			Price:          20000.00,
			PrincipalImage: "https://placehold.jp/3d4070/ffffff/150x150.png",
			OtherImages: []string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		}
		mockEntityProduct, _ := factory.NewProduct(
			mockProductPayload.Sku,
			mockProductPayload.Name,
			mockProductPayload.Brand,
			mockProductPayload.Size,
			mockProductPayload.Price,
			mockProductPayload.PrincipalImage,
			mockProductPayload.OtherImages,
		)

		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("CreateProduct", *mockEntityProduct).Return(errors.New("invalid sku format (right format: FAL-XXXXXXX)"))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.POST("/products", NewProductHandlers(mockUsecase).CreateProduct)
		payload, _ := json.Marshal(mockProductPayload)
		request, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(payload))

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(gin.H{
			"message": "invalid sku format (right format: FAL-XXXXXXX)",
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertNotCalled(t, "CreateProduct", mockEntityProduct)
	})
}

func TestGetProductBySku(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetProductBySku - 200 OK", func(t *testing.T) {

		sku := "FAL-1000000"
		mockEntityProduct, _ := factory.NewProduct(
			"FAL-1000000",
			"Polera",
			"CAT",
			"XL",
			20000.00,
			"https://placehold.jp/3d4070/ffffff/150x150.png",
			[]string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		)
		mockProductReturned := response.ConvertFromEntityToResponse(*mockEntityProduct)
		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("FindBySku", sku).Return(mockEntityProduct, nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.GET("/products/:sku", NewProductHandlers(mockUsecase).GetProductBySku)

		request, err := http.NewRequest(http.MethodGet, "/products/FAL-1000000", nil)

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(mockProductReturned)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertExpectations(t)
	})

	t.Run("GetProductBySku - 404 Not Found", func(t *testing.T) {

		sku := "FAL-9999999"
		resource := fmt.Sprintf("/products/%s", sku)
		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("FindBySku", sku).Return(nil, errors.New("record not found"))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.GET("/products/:sku", NewProductHandlers(mockUsecase).GetProductBySku)

		request, err := http.NewRequest(http.MethodGet, resource, nil)

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(gin.H{
			"message": "record not found",
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertExpectations(t)
	})
}

func TestGetAllProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAllProducts - 200 OK", func(t *testing.T) {
		mockEntityProduct1, _ := factory.NewProduct(
			"FAL-1000000",
			"Polera",
			"CAT",
			"XL",
			20000.00,
			"https://placehold.jp/3d4070/ffffff/150x150.png",
			[]string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		)

		mockEntityProduct2, _ := factory.NewProduct(
			"FAL-1000001",
			"Polera",
			"CAT",
			"L",
			15000.00,
			"https://placehold.jp/3d4070/ffffff/150x150.png",
			[]string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		)

		mockProductReturned1 := response.ConvertFromEntityToResponse(*mockEntityProduct1)
		mockProductReturned2 := response.ConvertFromEntityToResponse(*mockEntityProduct2)
		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("FindAll").Return(
			[]entity.Product{
				*mockEntityProduct1,
				*mockEntityProduct2,
			}, nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.GET("/products/", NewProductHandlers(mockUsecase).GetAllProducts)

		request, err := http.NewRequest(http.MethodGet, "/products/", nil)

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(
			[]response.ProductResponse{
				*mockProductReturned1,
				*mockProductReturned2,
			})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertExpectations(t)
	})

	t.Run("GetAllProducts - 404 Not found", func(t *testing.T) {
		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("FindAll").Return(nil, errors.New("records not found"))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.GET("/products/", NewProductHandlers(mockUsecase).GetAllProducts)

		request, err := http.NewRequest(http.MethodGet, "/products/", nil)

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(gin.H{
			"message": "records not found",
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("UpdateProduct - 200 Success", func(t *testing.T) {
		oldSku := "FAL-1000000"
		resource := fmt.Sprintf("/products/%s", oldSku)
		mockProductPayload := struct {
			Sku            string   `json:"sku"`
			Name           string   `json:"name"`
			Brand          string   `json:"brand"`
			Size           string   `json:"size"`
			Price          float64  `json:"price"`
			PrincipalImage string   `json:"principal_image"`
			OtherImages    []string `json:"other_images"`
		}{
			Sku:            "FAL-1000011",
			Name:           "Polera",
			Brand:          "CAT",
			Size:           "XL",
			Price:          20000.00,
			PrincipalImage: "https://placehold.jp/3d4070/ffffff/150x150.png",
			OtherImages: []string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		}
		mockEntityProduct, _ := factory.NewProduct(
			mockProductPayload.Sku,
			mockProductPayload.Name,
			mockProductPayload.Brand,
			mockProductPayload.Size,
			mockProductPayload.Price,
			mockProductPayload.PrincipalImage,
			mockProductPayload.OtherImages,
		)

		mockEntityProduct2, _ := factory.NewProduct(
			"FAL-1000000",
			"Polerón",
			"Ocean Pacific",
			"XL",
			25000.00,
			"https://placehold.jp/3d4070/ffffff/150x150.png",
			[]string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		)

		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("UpdateProduct", oldSku, *mockEntityProduct).Return(mockEntityProduct, nil)
		mockUsecase.On("FindBySku", oldSku).Return(mockEntityProduct2, nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.PUT("/products/:sku", NewProductHandlers(mockUsecase).UpdateProduct)
		payload, _ := json.Marshal(mockProductPayload)
		request, err := http.NewRequest(http.MethodPut, resource, bytes.NewBuffer(payload))

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(mockProductPayload)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertExpectations(t)
	})

	t.Run("UpdateProduct - 404 Not found", func(t *testing.T) {
		oldSku := "FAL-1000000"
		resource := fmt.Sprintf("/products/%s", oldSku)
		mockProductPayload := struct {
			Sku            string   `json:"sku"`
			Name           string   `json:"name"`
			Brand          string   `json:"brand"`
			Size           string   `json:"size"`
			Price          float64  `json:"price"`
			PrincipalImage string   `json:"principal_image"`
			OtherImages    []string `json:"other_images"`
		}{
			Sku:            "FAL-1000011",
			Name:           "Polera",
			Brand:          "CAT",
			Size:           "XL",
			Price:          20000.00,
			PrincipalImage: "https://placehold.jp/3d4070/ffffff/150x150.png",
			OtherImages: []string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		}
		mockEntityProduct, _ := factory.NewProduct(
			mockProductPayload.Sku,
			mockProductPayload.Name,
			mockProductPayload.Brand,
			mockProductPayload.Size,
			mockProductPayload.Price,
			mockProductPayload.PrincipalImage,
			mockProductPayload.OtherImages,
		)

		mockUsecase := new(usecase.UseCaseMock)
		notFoundError := errors.New("record not found")
		mockUsecase.On("UpdateProduct", oldSku, *mockEntityProduct).Return(nil, notFoundError)
		mockUsecase.On("FindBySku", oldSku).Return(nil, notFoundError)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.PUT("/products/:sku", NewProductHandlers(mockUsecase).UpdateProduct)
		payload, _ := json.Marshal(mockProductPayload)
		request, err := http.NewRequest(http.MethodPut, resource, bytes.NewBuffer(payload))

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(gin.H{
			"message": "record not found",
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())

		mockUsecase.AssertNotCalled(t, "UpdateProduct", oldSku, *mockEntityProduct)
	})

	t.Run("UpdateProduct - 422 Unprocessable Entity (Duplicated SKU)", func(t *testing.T) {
		oldSku := "FAL-1000000"
		resource := fmt.Sprintf("/products/%s", oldSku)
		mockProductPayload := struct {
			Sku            string   `json:"sku"`
			Name           string   `json:"name"`
			Brand          string   `json:"brand"`
			Size           string   `json:"size"`
			Price          float64  `json:"price"`
			PrincipalImage string   `json:"principal_image"`
			OtherImages    []string `json:"other_images"`
		}{
			Sku:            "FAL-1000011",
			Name:           "Polera",
			Brand:          "CAT",
			Size:           "XL",
			Price:          20000.00,
			PrincipalImage: "https://placehold.jp/3d4070/ffffff/150x150.png",
			OtherImages: []string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		}
		mockEntityProduct, _ := factory.NewProduct(
			mockProductPayload.Sku,
			mockProductPayload.Name,
			mockProductPayload.Brand,
			mockProductPayload.Size,
			mockProductPayload.Price,
			mockProductPayload.PrincipalImage,
			mockProductPayload.OtherImages,
		)

		mockEntityProduct2, _ := factory.NewProduct(
			"FAL-1000000",
			"Polerón",
			"Ocean Pacific",
			"XL",
			25000.00,
			"https://placehold.jp/3d4070/ffffff/150x150.png",
			[]string{
				"https://placehold.jp/30/dd6699/ffffff/300x150.png?text=placeholder+image",
				"https://placehold.jp/24/cccccc/ffffff/250x50.png?text=placehold.jp",
			},
		)

		mockUsecase := new(usecase.UseCaseMock)
		duplicatedSkuError := errors.New("duplicated sku")
		mockUsecase.On("UpdateProduct", oldSku, *mockEntityProduct).Return(nil, duplicatedSkuError)
		mockUsecase.On("FindBySku", oldSku).Return(mockEntityProduct2, nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.PUT("/products/:sku", NewProductHandlers(mockUsecase).UpdateProduct)
		payload, _ := json.Marshal(mockProductPayload)
		request, err := http.NewRequest(http.MethodPut, resource, bytes.NewBuffer(payload))

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response, err := json.Marshal(gin.H{
			"message": "duplicated sku",
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())

		// mockUsecase.AssertNotCalled(t, "UpdateProduct", oldSku, *mockEntityProduct)
		mockUsecase.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Delete - 204 No content", func(t *testing.T) {
		sku := "FAL-1000000"
		resource := fmt.Sprintf("/products/%s", sku)
		mockUsecase := new(usecase.UseCaseMock)

		mockUsecase.On("DeleteProduct", sku).Return(nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.Group("v1")
		router.DELETE("/products/:sku", NewProductHandlers(mockUsecase).Delete)

		request, err := http.NewRequest(http.MethodDelete, resource, nil)

		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		response := []byte(nil)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		assert.Equal(t, response, rr.Body.Bytes())
		mockUsecase.AssertExpectations(t)
	})

}
