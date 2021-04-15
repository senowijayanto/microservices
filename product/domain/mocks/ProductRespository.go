package mocks

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"product/domain"
)

type ProductRepository struct {
	mock.Mock
}

func (_m *ProductRepository) Fetch(ctx context.Context) ([]domain.Product, error)  {
	ret := _m.Called(ctx)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *ProductRepository) GetByID(ctx context.Context, id uint32) (domain.Product, error)  {
	ret := _m.Called(ctx, id)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, uint32) domain.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}