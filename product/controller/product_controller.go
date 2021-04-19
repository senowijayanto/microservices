package controller

import (
	"net/http"
	"product/domain"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProdService domain.ProductService
}

func NewProductController(e *echo.Echo, ps domain.ProductService) {
	controller := &ProductController{
		ProdService: ps,
	}
	group := e.Group("/api/v1")
	group.GET("/products", controller.Fetch)
	group.GET("/products/:id", controller.GetByID)
	group.POST("/products", controller.Store)
	group.PUT("/products/:id", controller.Update)
	group.DELETE("/products/:id", controller.Delete)
	group.POST("/products/order/:id", controller.Order)
}

func (ph *ProductController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	list, err := ph.ProdService.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, list)
}

func (ph *ProductController) GetByID(c echo.Context) error {
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

func (ph *ProductController) Store(c echo.Context) (err error) {
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

func (ph *ProductController) Update(c echo.Context) error {
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

func (ph *ProductController) Delete(c echo.Context) error {
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

func (ph *ProductController) Order(c echo.Context) (err error) {
	var po domain.ProductOrder
	err = c.Bind(&po)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": err.Error(),
		})
	}

	// Get product by ID
	paramID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err.Error(),
		})
	}

	id := uint32(paramID)
	ctx := c.Request().Context()
	qty := po.Qty
	product, err := ph.ProdService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": "Record Not Found",
		})
	}

	// Check stock
	stock := product.Stock
	stockNow := stock - qty
	if stockNow < 0 {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": "Insufficient stock",
		})
	}

	// Update stock
	product.Stock = stockNow
	product.UpdatedAt = time.Now()
	err = ph.ProdService.UpdateStock(ctx, &product, id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)

	// TODO
	// return c.JSON(http.StatusOK, echo.Map{"message": "Order Accepted"})

	// fmt.Println("Consumer Application")
	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// defer conn.Close()

	// ch, err := conn.Channel()
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// defer ch.Close()

	// msgs, err := ch.Consume(
	// 	"OrderQueue",
	// 	"",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )

	// forever := make(chan bool)
	// go func() {
	// 	for d := range msgs {
	// 		fmt.Printf("Received Message: %s\n", d.Body)
	// 	}
	// }()

	// fmt.Println("Successfully connected to our RabbitMQ Instance")
	// fmt.Println(" [+] - waiting for messages")
	// <-forever
	// return

}
