package domain

import (
	"context"
	"time"
)

type Product struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductOrder struct {
	ID  uint32 `json:"id"`
	Qty int    `json:"qty"`
}

type ProductRepository interface {
	Fetch(ctx context.Context) (products []Product, err error)
	GetByID(ctx context.Context, id uint32) (product Product, err error)
	Store(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product, id uint32) error
	Delete(ctx context.Context, id uint32) error
	UpdateStock(ctx context.Context, product *Product, id uint32) error
}

type ProductService interface {
	Fetch(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id uint32) (Product, error)
	Store(context.Context, *Product) error
	Update(ctx context.Context, product *Product, id uint32) error
	Delete(ctx context.Context, id uint32) error
	UpdateStock(ctx context.Context, product *Product, id uint32) error
}
