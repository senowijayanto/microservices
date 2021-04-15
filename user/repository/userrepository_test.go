package repository

import (
	"context"
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
	"user/domain"
)

var user = &domain.User{
	ID:        1,
	Username:  "senowijayanto",
	Email:     "senowijayanto@gmail.com",
	Password:  "k03nc1k03",
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

func TestUserRepository_Fetch(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, username, email, created_at, updated_at FROM user"

	rows := sqlmock.NewRows([]string{"id", "username", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.Fetch(context.TODO())
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestUserRepository_GetByID(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	query := "SELECT id, username, email, created_at, updated_at FROM user WHERE id=\\?"

	rows := sqlmock.NewRows([]string{"id", "username", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(user.ID).WillReturnRows(rows)

	u, err := repo.GetByID(context.TODO(), user.ID)
	assert.NotNil(t, u)
	assert.NoError(t, err)

}

func TestUserRepository_Store(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	hashing, err := HashPassword(user.Password)
	if err != nil {
		return
	}

	user.Password = hashing

	query := "INSERT INTO user \\(username, email, password, created_at, updated_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Store(context.TODO(), user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_Update(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer func() {
		db.Close()
	}()

	hashing, err := HashPassword(user.Password)
	if err != nil {
		return
	}

	query := "UPDATE user SET username=\\?, email=\\?, password=\\?, updated_at=\\? WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.Username, user.Email, hashing, time.Now(), user.ID).WillReturnResult(sqlmock.NewResult(1,1))

	err = repo.Update(context.TODO(), user, uint32(1))
	assert.NoError(t, err)
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
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