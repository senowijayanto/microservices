package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"user/domain"
)

// UserController represent the http handler for user
type UserController struct {
	UserService domain.UserService
}

// NewUserController will initialize the users/resources endpoint
func NewUserController(e *echo.Echo, us domain.UserService) {
	controller := &UserController{
		UserService: us,
	}
	group := e.Group("/api/v1")
	group.GET("/users", controller.Fetch)
	group.GET("/users/:id", controller.GetByID)
	group.POST("/users", controller.Store)
	group.PUT("/users/:id", controller.Update)
	group.DELETE("/users/:id", controller.Delete)
}

// Fetch method will fetch all users data
func (uc *UserController) Fetch(c echo.Context) error  {
	ctx := c.Request().Context()

	users, err := uc.UserService.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetByID(c echo.Context) error  {
	paramID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err.Error(),
		})
	}

	id := uint32(paramID)
	ctx := c.Request().Context()

	user, err := uc.UserService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) Store(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": err.Error(),
		})
	}

	ctx := c.Request().Context()
	err = uc.UserService.Store(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"email": user.Email,
	})
}

func (uc *UserController) Update(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": err.Error(),
		})
	}

	paramID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err.Error(),
		})
	}

	id := uint32(paramID)
	ctx := c.Request().Context()

	err = uc.UserService.Update(ctx, &user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) Delete(c echo.Context) error  {
	paramID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}

	id := uint32(paramID)
	ctx := c.Request().Context()

	err = uc.UserService.Delete(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}