package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"product/domain"
	"product/domain/mocks"
	"testing"
	"time"
)

func TestProductService_Fetch(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{
		Name:      "Laptop Lenovo",
		Price:     3000000,
	}

	mockListProduct := make([]domain.Product, 0)
	mockListProduct = append(mockListProduct, mockProduct)

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Fetch", mock.Anything).Return(mockListProduct, nil).Once()
		p := NewProductService(mockProductRepo, time.Second*2)
		list, err := p.Fetch(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListProduct))

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("Fetch", mock.Anything).Return(nil, errors.New("unexpected error")).Once()
		p := NewProductService(mockProductRepo, time.Second*2)
		list, err := p.Fetch(context.TODO())

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductService_GetByID(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{Name: "Laptop Lenovo", Price: 3000000}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("uint32")).Return(mockProduct, nil).Once()
		p := NewProductService(mockProductRepo, time.Second*2)

		res, err := p.GetByID(context.TODO(), mockProduct.ID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("uint32")).Return(domain.Product{}, errors.New("unexpected error")).Once()
		p := NewProductService(mockProductRepo, time.Second*2)

		res, err := p.GetByID(context.TODO(), mockProduct.ID)
		assert.Error(t, err)
		assert.Equal(t, domain.Product{}, res)
		mockProductRepo.AssertExpectations(t)
	})
}