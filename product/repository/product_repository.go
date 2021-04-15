package repository

import (
	"context"
	"database/sql"
	"log"
	"product/domain"
	"time"
)

type productRepository struct {
	Conn *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		Conn: db,
	}
}

func (pr *productRepository) Fetch(ctx context.Context) (products []domain.Product, err error)  {
	query := `SELECT id, name, price, quantity, created_at, updated_at FROM product`
	rows, err := pr.Conn.QueryContext(ctx, query)
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

	products = make([]domain.Product, 0)
	for rows.Next() {
		p := domain.Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, p)
	}

	return
}

func (pr *productRepository) GetByID(ctx context.Context, id uint32) (product domain.Product, err error)  {
	query := `SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id=?`

	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return domain.Product{}, err
	}

	row := stmt.QueryRowContext(ctx, id)
	product = domain.Product{}

	err = row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.UpdatedAt)
	return
}

func (pr *productRepository) Store(ctx context.Context, product *domain.Product) (err error)  {
	query := `INSERT INTO product (name, price, quantity, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Quantity, time.Now(), time.Now())
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	product.ID = uint32(lastID)
	return
}