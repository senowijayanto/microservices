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
