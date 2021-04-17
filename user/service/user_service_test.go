package service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"user/domain"
	"user/domain/mocks"
	"user/service"
)

var (
	mockUserRepo = new(mocks.UserRepository)
	mockUser     = domain.User{Email: "senowijayanto@gmail.com"}
	u            = service.NewUserService(mockUserRepo, time.Second*2)
)

func TestUserService_Fetch(t *testing.T) {
	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Fetch", mock.Anything).Return(mockListUser, nil).Once()
		list, err := u.Fetch(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListUser))

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("Fetch", mock.Anything).Return(nil, errors.New("unexpected error")).Once()
		list, err := u.Fetch(context.TODO())

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("uint32")).Return(mockUser, nil).Once()

		res, err := u.GetByID(context.TODO(), mockUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("uint32")).Return(domain.User{}, errors.New("unexpected error")).Once()

		res, err := u.GetByID(context.TODO(), mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, domain.User{}, res)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_Store(t *testing.T) {
	tempMockUser := mockUser
	tempMockUser.ID = 0

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := u.Store(context.TODO(), &tempMockUser)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Email, tempMockUser.Email)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Update", mock.Anything, &mockUser).Once().Return(nil)

		err := u.Update(context.TODO(), &mockUser, mockUser.ID)
		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Delete", mock.Anything, mock.AnythingOfType("uint32")).Return(nil).Once()

		err := u.Delete(context.TODO(), mockUser.ID)
		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}
