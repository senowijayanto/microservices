package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product/domain"
	"strconv"
)

type ProductHandler struct {
	ProdService domain.ProductService
}

func NewProductHandler(e *echo.Echo, ps domain.ProductService)  {
	handler := &ProductHandler{
		ProdService: ps,
	}
	e.GET("/api/v1/products", handler.Fetch)
	e.GET("/api/v1/products/:id", handler.GetByID)
	e.POST("/api/v1/products", handler.Store)
	e.PUT("/api/v1/products/:id", handler.Update)
	e.DELETE("/api/v1/products/:id", handler.Delete)
}

func (ph *ProductHandler) Fetch(c echo.Context) error  {
	ctx := c.Request().Context()

	list, err := ph.ProdService.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, list)
}

func (ph *ProductHandler) GetByID(c echo.Context) error  {
	paramID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}

	id := uint32(paramID)
	ctx := c.Request().Context()

	product, err := ph.ProdService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) Store(c echo.Context) (err error) {
	var product domain.Product
	err = c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": err.Error(),
		})
	}

	ctx := c.Request().Context()
	err = ph.ProdService.Store(ctx, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"name": product.Name,
	})
}

func (ph *ProductHandler) Update(c echo.Context) error {
	var product domain.Product
	err := c.Bind(&product)
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

	err = ph.ProdService.Update(ctx, &product, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	paramID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err.Error(),
		})
	}

	id := uint32(paramID)
	ctx := c.Request().Context()

	err = ph.ProdService.Delete(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}


