package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"product/controller"
	"product/domain"
	"product/domain/mocks"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProductController_Fetch(t *testing.T) {
	var mockProduct domain.Product
	mockProdService := new(mocks.ProductService)
	mockListProduct := make([]domain.Product, 0)
	mockListProduct = append(mockListProduct, mockProduct)

	mockProdService.On("Fetch", mock.Anything).Return(mockListProduct, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/products", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := controller.ProductController{ProdService: mockProdService}
	err = handler.Fetch(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockProdService.AssertExpectations(t)
}

func TestProductController_Fetch_Error(t *testing.T) {
	mockProdService := new(mocks.ProductService)
	mockProdService.On("Fetch", mock.Anything).Return(nil, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/products", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := controller.ProductController{ProdService: mockProdService}
	err = handler.Fetch(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProdService.AssertExpectations(t)
}

func TestProductController_GetByID(t *testing.T) {
	var mockProduct domain.Product

	mockProdService := new(mocks.ProductService)
	num := int(mockProduct.ID)
	mockProdService.On("GetByID", mock.Anything, uint32(num)).Return(mockProduct, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/products/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/products/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := controller.ProductController{ProdService: mockProdService}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProdService.AssertExpectations(t)
}

func TestProductController_Store(t *testing.T) {
	mockProduct := domain.Product{
		ID:        1,
		Name:      "Laptop Lenovo Thinkpad",
		Price:     7000000,
		Stock:     10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockProduct := mockProduct
	tempMockProduct.ID = 0
	mockProdService := new(mocks.ProductService)

	jm, err := json.Marshal(tempMockProduct)
	assert.NoError(t, err)

	mockProdService.On("Store", mock.Anything, mock.AnythingOfType("*domain.Product")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/api/v1/products", strings.NewReader(string(jm)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/products")

	handler := controller.ProductController{ProdService: mockProdService}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockProdService.AssertExpectations(t)
}

func TestProductController_Update(t *testing.T) {
	mockProduct := domain.Product{
		ID:        1,
		Name:      "Laptop Lenovo Thinkpad",
		Price:     7000000,
		Stock:     10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	tempMockProduct := mockProduct
	tempMockProduct.ID = 0
	mockProdService := new(mocks.ProductService)

	jm, err := json.Marshal(tempMockProduct)
	assert.NoError(t, err)

	num := int(mockProduct.ID)

	mockProdService.On("Update", mock.Anything, mock.AnythingOfType("*domain.Product"), uint32(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/api/v1/products/"+strconv.Itoa(num), strings.NewReader(string(jm)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/products/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))

	handler := controller.ProductController{ProdService: mockProdService}
	err = handler.Update(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockProdService.AssertExpectations(t)
}

func TestProductController_Delete(t *testing.T) {
	var mockProduct domain.Product
	mockProdService := new(mocks.ProductService)
	num := int(mockProduct.ID)

	mockProdService.On("Delete", mock.Anything, uint32(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/api/v1/products/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/products/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))

	handler := controller.ProductController{ProdService: mockProdService}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockProdService.AssertExpectations(t)
}
