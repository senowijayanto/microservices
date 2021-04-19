package domain

import (
	"golang.org/x/net/context"
	"time"
)

type Order struct {
	ID        uint32    `json:"id"`
	ProductID uint32    `json:"product_id"`
	UserID    uint32    `json:"user_id"`
	Qty       int       `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderRepository interface {
	Fetch(ctx context.Context) (orders []Order, err error)
	Store(ctx context.Context, order *Order) error
}

type OrderService interface {
	Fetch(ctx context.Context) ([]Order, error)
	Store(context.Context, *Order) error
}
