package repository_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"order/domain"
	"order/repository"
	"regexp"
	"testing"
	"time"
)

var (
	t       = time.Now()
	ts      = t.Format("2006-01-02 15:04:05")
	order = &domain.Order{
		ID: 1,
		ProductID: 3,
		Qty: 2,
	}
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestOrderRepository_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := regexp.QuoteMeta(`INSERT INTO order (product_id, user_id, qty, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(order.ProductID, order.Qty, ts, ts).WillReturnResult(sqlmock.NewResult(1, 1))

	or := repository.NewOrderRepository(db)

	err = or.Store(context.TODO(), order)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), order.ID)
}
