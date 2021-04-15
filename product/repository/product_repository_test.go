package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"product/domain"
	"testing"
	"time"
)

var product = &domain.Product{
	ID:        1,
	Name:      "Laptop Lenovo",
	Price:     3000000,
	Quantity:  10,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestProductRepository_Fetch(t *testing.T) {
	db, mock := NewMock()
	repo := NewProductRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, name, price, quantity, created_at, updated_at FROM product"

	rows := sqlmock.NewRows([]string{"id", "name", "price", "quantity", "created_at", "updated_at"}).
		AddRow(product.ID, product.Name, product.Price, product.Quantity, product.CreatedAt, product.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	prod, err := repo.Fetch(context.TODO())
	assert.NotEmpty(t, prod)
	assert.NoError(t, err)
	assert.Len(t, prod, 1)
}

func TestProductRepository_GetByID(t *testing.T) {
	db, mock := NewMock()
	repo := NewProductRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, name, price, quantity, created_at, updated_at FROM product WHERE id=\\?"

	rows := sqlmock.NewRows([]string{"id", "name", "price", "quantity", "created_at", "updated_at"}).
		AddRow(product.ID, product.Name, product.Price, product.Quantity, product.CreatedAt, product.UpdatedAt)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(product.ID).WillReturnRows(rows)

	p, err := repo.GetByID(context.TODO(), product.ID)
	assert.NotNil(t, p)
	assert.NoError(t, err)
}

func TestProductRepository_Store(t *testing.T) {
	now := time.Now()
	prod := &domain.Product{
		ID:        1,
		Name:      "Laptop Lenovo",
		Price:     3000000,
		Quantity:  10,
		CreatedAt: now,
		UpdatedAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT INTO product \\(name, price, quantity, created_at, updated_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(prod.Name, prod.Price, prod.Quantity, prod.CreatedAt, prod.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	pr := NewProductRepository(db)

	err = pr.Store(context.TODO(), prod)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), prod.ID)
}