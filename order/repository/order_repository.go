package repository

import (
	"database/sql"
	"golang.org/x/net/context"
	"log"
	"order/domain"
	"time"
)

var (
	t  = time.Now()
	ts = t.Format("2006-01-02 15:04:05")
)

type orderRepository struct {
	Conn *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository  {
	return &orderRepository{Conn: db}
}

func (or *orderRepository) Fetch(ctx context.Context) (orders []domain.Order, err error)  {
	query := "SELECT id, product_id, user_id, qty, created_at, updated_at FROM `order`"
	rows, err := or.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Fatal(errRow)
		}
	}()

	orders = make([]domain.Order, 0)
	for rows.Next() {
		o := domain.Order{}
		err = rows.Scan(&o.ID, &o.ProductID, &o.UserID, &o.Qty, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		orders = append(orders, o)
	}

	return
}

func (or *orderRepository) Store(ctx context.Context, order *domain.Order) (err error)  {
	query := "INSERT INTO `order` (product_id, user_id, qty, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := or.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, order.ProductID, order.UserID, order.Qty, ts, ts)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	order.ID = uint32(lastID)
	return
}
