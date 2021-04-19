package service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"product/domain"
	"product/domain/mocks"
	"product/service"
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
		p := service.NewProductService(mockProductRepo, time.Second*2)
		list, err := p.Fetch(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListProduct))

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("Fetch", mock.Anything).Return(nil, errors.New("unexpected error")).Once()
		p := service.NewProductService(mockProductRepo, time.Second*2)
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
		p := service.NewProductService(mockProductRepo, time.Second*2)

		res, err := p.GetByID(context.TODO(), mockProduct.ID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("uint32")).Return(domain.Product{}, errors.New("unexpected error")).Once()
		p := service.NewProductService(mockProductRepo, time.Second*2)

		res, err := p.GetByID(context.TODO(), mockProduct.ID)
		assert.Error(t, err)
		assert.Equal(t, domain.Product{}, res)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductService_Store(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{Name: "Laptop Lenovo", Price: 3000000}
	tempMockProduct := mockProduct
	tempMockProduct.ID = 0

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Product")).Return(nil).Once()

		p := service.NewProductService(mockProductRepo, time.Second*2)
		err := p.Store(context.TODO(), &tempMockProduct)

		assert.NoError(t, err)
		assert.Equal(t, mockProduct.Name, tempMockProduct.Name)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductService_Update(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{ID: 1, Name: "Laptop Lenovo", Price: 3000000}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Update", mock.Anything, &mockProduct).Once().Return(nil)
		p := service.NewProductService(mockProductRepo, time.Second*2)

		err := p.Update(context.TODO(), &mockProduct, mockProduct.ID)
		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductService_Delete(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{ID: 1, Name: "Laptop Lenovo", Price: 3000000}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Delete", mock.Anything, mock.AnythingOfType("uint32")).Return(nil).Once()
		p := service.NewProductService(mockProductRepo, time.Second*2)

		err := p.Delete(context.TODO(), mockProduct.ID)
		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductService_UpdateStock(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{ID: 1, Stock: 15}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Update", mock.Anything, &mockProduct).Once().Return(nil)
		p := service.NewProductService(mockProductRepo, time.Second*2)

		err := p.UpdateStock(context.TODO(), &mockProduct, mockProduct.ID)
		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}