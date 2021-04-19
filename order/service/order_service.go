package service

import (
	"golang.org/x/net/context"
	"order/domain"
	"time"
)

type orderService struct {
	orderRepo domain.OrderRepository
	contextTimeout time.Duration
}

func NewOrderService(order domain.OrderRepository, timeout time.Duration) domain.OrderService {
	return &orderService{
		orderRepo:      order,
		contextTimeout: timeout,
	}
}

func (os *orderService) Fetch(c context.Context) (orders []domain.Order, err error) {
	ctx, cancel := context.WithTimeout(c, os.contextTimeout)
	defer cancel()

	orders, err = os.orderRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (os *orderService) Store(c context.Context, order *domain.Order) (err error)  {
	ctx, cancel := context.WithTimeout(c, os.contextTimeout)
	defer cancel()

	err = os.orderRepo.Store(ctx, order)
	return
}