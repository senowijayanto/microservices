package service

import (
	"context"
	"product/domain"
	"time"
)

type productService struct {
	productRepo domain.ProductRepository
	contextTimeout time.Duration
}

func NewProductService(product domain.ProductRepository, timeout time.Duration) domain.ProductService  {
	return &productService{
		productRepo: product,
		contextTimeout: timeout,
	}
}

func (ps *productService) Fetch(c context.Context) (products []domain.Product, err error)  {
	ctx, cancel := context.WithTimeout(c, ps.contextTimeout)
	defer cancel()

	products, err = ps.productRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (ps *productService) GetByID(c context.Context, id uint32) (product domain.Product, err error)  {
	ctx, cancel := context.WithTimeout(c, ps.contextTimeout)
	defer cancel()

	product, err = ps.productRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}

func (ps *productService) Store(c context.Context, product *domain.Product) (err error)  {
	ctx, cancel := context.WithTimeout(c, ps.contextTimeout)
	defer cancel()

	err = ps.productRepo.Store(ctx, product)
	return 
}

func (ps *productService) Update(c context.Context, product *domain.Product, id uint32) (err error)  {
	ctx, cancel := context.WithTimeout(c, ps.contextTimeout)
	defer cancel()

	product.UpdatedAt = time.Now()
	return ps.productRepo.Update(ctx, product, id)
}

func (ps *productService) Delete(c context.Context, id uint32) (err error)  {
	ctx, cancel := context.WithTimeout(c, ps.contextTimeout)
	defer cancel()

	err = ps.productRepo.Delete(ctx, id)
	return
}

func (ps productService) UpdateStock(c context.Context, product *domain.Product, id uint32) (err error)  {
	ctx, cancel := context.WithTimeout(c, ps.contextTimeout)
	defer cancel()

	product.UpdatedAt = time.Now()
	return ps.productRepo.Update(ctx, product, id)
}