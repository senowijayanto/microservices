package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"product/domain"
	"time"
)

var (
	t  = time.Now()
	ts = t.Format("2006-01-02 15:04:05")
)

type productRepository struct {
	Conn *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		Conn: db,
	}
}

func (pr *productRepository) Fetch(ctx context.Context) (products []domain.Product, err error) {
	query := `SELECT id, name, price, stock, created_at, updated_at FROM product`
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
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, p)
	}

	return
}

func (pr *productRepository) GetByID(ctx context.Context, id uint32) (product domain.Product, err error) {
	query := `SELECT id, name, price, stock, created_at, updated_at FROM product WHERE id=?`

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
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt)
	return
}

func (pr *productRepository) Store(ctx context.Context, product *domain.Product) (err error) {
	query := `INSERT INTO product (name, price, stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Stock, ts, ts)
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

func (pr *productRepository) Update(ctx context.Context, product *domain.Product, id uint32) (err error) {
	query := `UPDATE product SET name=?, price=?, stock=?, updated_at=? WHERE id=?`

	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	result, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Stock, ts, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("total affected: %d", rowsAffected)
		return
	}

	return
}

func (pr *productRepository) Delete(ctx context.Context, id uint32) (err error) {
	query := `DELETE FROM product WHERE id=?`

	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("total affected: %d", rowsAffected)
		return
	}

	return
}

func (pr *productRepository) UpdateStock(ctx context.Context, product *domain.Product, id uint32) (err error) {
	query := `UPDATE product SET stock=?, updated_at=? WHERE id=?`

	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	t := time.Now()
	ts := t.Format("2006-01-02 15:04:05")
	result, err := stmt.ExecContext(ctx, product.Stock, ts, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("total affected: %d", rowsAffected)
		return
	}

	return
}
