package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint32    `json:"id"`
	Email     string    `json:"email" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserRepository interface {
	Fetch(ctx context.Context) (users []User, err error)
	GetByID(ctx context.Context, id uint32) (user User, err error)
	Store(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User, id uint32) error
	Delete(ctx context.Context, id uint32) error
}

type UserService interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uint32) (User, error)
	Store(context.Context, *User) error
	Update(ctx context.Context, user *User, id uint32) error
	Delete(ctx context.Context, id uint32) error
}