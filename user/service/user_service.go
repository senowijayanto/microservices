package service

import (
	"context"
	"time"
	"user/domain"
)

type userService struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserService will create new an userService object representation of domain.UserService interface
func NewUserService(user domain.UserRepository, timeout time.Duration) domain.UserService {
	return &userService{
		userRepo: user,
		contextTimeout: timeout,
	}
}

func (us *userService) Fetch(c context.Context) (users []domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	users, err = us.userRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func (us *userService) GetByID(c context.Context, id uint32) (user domain.User, err error)  {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	user, err = us.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}

func (us *userService) Store(c context.Context, user *domain.User) (err error)  {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	err = us.userRepo.Store(ctx, user)
	return
}

func (us *userService) Update(c context.Context, user *domain.User, id uint32) (err error)  {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	user.UpdatedAt = time.Now()
	return us.userRepo.Update(ctx, user, id)
}

func (us *userService) Delete(c context.Context, id uint32) (err error)  {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	err = us.userRepo.Delete(ctx, id)
	return
}