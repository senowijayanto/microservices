package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"product/domain"
	"product/domain/mocks"
	"strconv"
	"strings"
	"testing"
)

func TestProductHandler_Fetch(t *testing.T) {
	var mockProduct domain.Product
	mockProdService := new(mocks.ProductService)
	mockListProduct := make([]domain.Product, 0)
	mockListProduct = append(mockListProduct, mockProduct)

	mockProdService.On("Fetch", mock.Anything).Return(mockListProduct, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := ProductHandler{ProdService: mockProdService}
	err = handler.Fetch(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockProdService.AssertExpectations(t)
}

func TestProductHandler_Fetch_Error(t *testing.T) {
	mockProdService := new(mocks.ProductService)
	mockProdService.On("Fetch", mock.Anything).Return(nil, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/products", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := ProductHandler{ProdService: mockProdService}
	err = handler.Fetch(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProdService.AssertExpectations(t)
}

func TestProductHandler_GetByID(t *testing.T) {
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
	handler := ProductHandler{ProdService: mockProdService}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProdService.AssertExpectations(t)
}