package repository

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"testing"
	"time"
	"user/domain"
)

func TestUserRepository_Fetch(t *testing.T) {
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	mockUsers := []domain.User{
		{
			ID:        1,
			Username:  "andrew",
			Email:     "andrew@gmail.com",
			Password:  "andrew1234",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Username:  "john",
			Email:     "john@gmail.com",
			Password:  "john1234",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "created_at", "updated_at"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Username, mockUsers[0].Email, mockUsers[0].Password, mockUsers[0].CreatedAt, mockUsers[0].UpdatedAt).
		AddRow(mockUsers[1].ID, mockUsers[1].Username, mockUsers[1].Email, mockUsers[1].Password, mockUsers[1].CreatedAt, mockUsers[1].UpdatedAt)

	query := "SELECT id, username, email, password, created_at, updated_at FROM user"

	mock.ExpectQuery(query).WillReturnRows(rows)
}

func TestUserRepository_GetByID(t *testing.T) {
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "created_at", "updated_at"}).
		AddRow(1, "andrew", "andrew@gmail.com", "andrew1234", time.Now(), time.Now())

	query := "SELECT id, username, email, password, created_at, updated_at FROM user"

	mock.ExpectQuery(query).WillReturnRows(rows)
}

func TestUserRepository_Store(t *testing.T) {
	now := time.Now()
	user := &domain.User{
		Username:  "andrew",
		Email:     "andrew@gmail.com",
		Password:  "andrew1234",
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT INTO user (username, email, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))
}