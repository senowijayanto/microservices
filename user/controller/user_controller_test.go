package controller_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
	"user/controller"
	"user/domain"
	"user/domain/mocks"
)

var (
	mockUserService = new(mocks.UserService)
)

func TestUserController_Fetch(t *testing.T) {
	var mockUser domain.User
	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)

	mockUserService.On("Fetch", mock.Anything).Return(mockListUser, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := controller.UserController{UserService: mockUserService}
	err = handler.Fetch(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserController_Fetch_Error(t *testing.T) {
	mockUserService.On("Fetch", mock.Anything).Return(nil, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := controller.UserController{UserService: mockUserService}
	err = handler.Fetch(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserController_GetByID(t *testing.T) {
	var mockUser domain.User

	num := int(mockUser.ID)
	mockUserService.On("GetByID", mock.Anything, uint32(num)).Return(mockUser, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/users/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := controller.UserController{UserService: mockUserService}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserController_Store(t *testing.T) {
	mockUser := domain.User{
		Email:     "senowijayanto@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockUser := mockUser
	tempMockUser.ID = 0

	jm, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	mockUserService.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/api/v1/products", strings.NewReader(string(jm)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users")

	handler := controller.UserController{UserService: mockUserService}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserController_Update(t *testing.T) {
	mockUser := domain.User{
		Email:     "senowijayanto@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockUser := mockUser
	tempMockUser.ID = 0

	jm, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	num := int(mockUser.ID)

	mockUserService.On("Update", mock.Anything, mock.AnythingOfType("*domain.User"), uint32(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/api/v1/users/"+strconv.Itoa(num), strings.NewReader(string(jm)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))

	handler := controller.UserController{UserService: mockUserService}
	err = handler.Update(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserController_Delete(t *testing.T) {
	var mockUser domain.User
	num := int(mockUser.ID)

	mockUserService.On("Delete", mock.Anything, uint32(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/api/v1/users/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))

	handler := controller.UserController{UserService: mockUserService}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUserService.AssertExpectations(t)
}