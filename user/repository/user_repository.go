package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"user/domain"
)

var (
	t = time.Now()
	ts = t.Format("2006-01-02 15:04:05")
)

type userRepository struct {
	Conn *sql.DB
}

// NewUserRepository will create an object that represent the user.Repository interface
func NewUserRepository(Conn *sql.DB) domain.UserRepository {
	return &userRepository{Conn}
}

func (ur *userRepository) Fetch(ctx context.Context) (users []domain.User, err error)  {
	query := `SELECT id, email, created_at, updated_at FROM user`
	rows, err := ur.Conn.QueryContext(ctx, query)
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

	users = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(&t.ID, &t.Email, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users,t)
	}

	return
}

func (ur *userRepository) GetByID(ctx context.Context, id uint32) (user domain.User, err error)  {
	query := `SELECT id, email, created_at, updated_at FROM user WHERE id=?`

	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return domain.User{}, err
	}
	row := stmt.QueryRowContext(ctx, id)
	user = domain.User{}

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		)
	return
}

func (ur *userRepository) Store(ctx context.Context, user *domain.User) (err error)  {
	query := `INSERT INTO user (email, created_at, updated_at) VALUES (?, ?, ?)`
	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, user.Email, ts, ts)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	user.ID = uint32(lastID)
	return
}

func (ur *userRepository) Update(ctx context.Context, user *domain.User, id uint32) (err error)  {
	query := `UPDATE user SET email=?, updated_at=? WHERE id=?`

	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}


	res, err := stmt.ExecContext(ctx, user.Email, ts, id)
	if err != nil {
		return
	}

	rowsEffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsEffected != 1 {
		err = fmt.Errorf("Total Affected: %d", rowsEffected)
		return
	}

	return
}

func (ur *userRepository) Delete(ctx context.Context, id uint32) (err error)  {
	query := `DELETE FROM user WHERE id=?`

	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsEffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsEffected != 1 {
		err = fmt.Errorf("Total Affected: %d", rowsEffected)
		return
	}

	return
}