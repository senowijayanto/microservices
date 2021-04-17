package repository_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
	"time"
	"user/domain"
	"user/repository"
)
var (
	t    = time.Now()
	ts   = t.Format("2006-01-02 15:04:05")
	user = &domain.User{
		ID:    1,
		Email: "senowijayanto@gmail.com",
	}
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestUserRepository_Fetch(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, email, created_at, updated_at FROM user"

	rows := sqlmock.NewRows([]string{"id", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Email, user.CreatedAt, user.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.Fetch(context.TODO())
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestUserRepository_GetByID(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, email, created_at, updated_at FROM user WHERE id=\\?"

	rows := sqlmock.NewRows([]string{"id", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Email, user.CreatedAt, user.UpdatedAt)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(user.ID).WillReturnRows(rows)

	u, err := repo.GetByID(context.TODO(), user.ID)
	assert.NotNil(t, u)
	assert.NoError(t, err)

}

func TestUserRepository_Store(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	query := regexp.QuoteMeta("INSERT INTO user (email, created_at, updated_at) VALUES (?, ?, ?)")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(user.Email, ts, ts).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Store(context.TODO(), user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_Update(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	query := regexp.QuoteMeta("UPDATE user SET email=?, updated_at=? WHERE id=?")

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.Email, ts, user.ID).WillReturnResult(sqlmock.NewResult(1,1))

	err := repo.Update(context.TODO(), user, uint32(1))
	assert.NoError(t, err)
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	query := "DELETE FROM user WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	num := uint32(1)
	err := repo.Delete(context.TODO(), num)
	assert.NoError(t, err)
}