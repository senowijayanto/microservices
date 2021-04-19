package repository_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"product/domain"
	"product/repository"
	"regexp"
	"testing"
	"time"
)

var (
	t       = time.Now()
	ts      = t.Format("2006-01-02 15:04:05")
	product = &domain.Product{
		ID:    1,
		Name:  "Laptop Lenovo",
		Price: 3000000,
		Stock: 10,
	}
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestProductRepository_Fetch(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewProductRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, name, price, quantity, created_at, updated_at FROM product"

	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock", "created_at", "updated_at"}).
		AddRow(product.ID, product.Name, product.Price, product.Stock, product.CreatedAt, product.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	prod, err := repo.Fetch(context.TODO())
	assert.NotEmpty(t, prod)
	assert.NoError(t, err)
	assert.Len(t, prod, 1)
}

func TestProductRepository_GetByID(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewProductRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id=\\?"

	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock", "created_at", "updated_at"}).
		AddRow(product.ID, product.Name, product.Price, product.Stock, product.CreatedAt, product.UpdatedAt)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(product.ID).WillReturnRows(rows)

	p, err := repo.GetByID(context.TODO(), product.ID)
	assert.NotNil(t, p)
	assert.NoError(t, err)
}

func TestProductRepository_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := regexp.QuoteMeta(`INSERT INTO product (name, price, stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(product.Name, product.Price, product.Stock, ts, ts).WillReturnResult(sqlmock.NewResult(1, 1))

	pr := repository.NewProductRepository(db)

	err = pr.Store(context.TODO(), product)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), product.ID)
}

func TestProductRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := regexp.QuoteMeta(`UPDATE product SET name=?, price=?, stock=?, updated_at=? WHERE id=?`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(product.Name, product.Price, product.Stock, ts, product.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	pr := repository.NewProductRepository(db)

	err = pr.Update(context.TODO(), product, product.ID)
	assert.NoError(t, err)
}

func TestProductRepository_Delete(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewProductRepository(db)
	defer func() {
		db.Close()
	}()

	query := regexp.QuoteMeta(`DELETE FROM product WHERE id=?`)

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	num := uint32(1)
	err := repo.Delete(context.TODO(), num)
	assert.NoError(t, err)
}

func TestProductRepository_UpdateStock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := regexp.QuoteMeta(`UPDATE product SET stock=?, updated_at=? WHERE id=?`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(product.Stock, ts, product.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	pr := repository.NewProductRepository(db)

	err = pr.UpdateStock(context.TODO(), product, product.ID)
	assert.NoError(t, err)
}
