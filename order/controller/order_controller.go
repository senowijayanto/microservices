package controller

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"order/domain"
	"strconv"
)

type OrderController struct {
	OrderService domain.OrderService
}

func NewOrderController(e *echo.Echo, os domain.OrderService) {
	controller := &OrderController{OrderService: os}

	group := e.Group("/api/v1")
	group.GET("/orders", controller.Fetch)
	group.POST("/orders", controller.Store)
}

func (oc *OrderController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	list, err := oc.OrderService.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, list)
}

func (oc *OrderController) Store(c echo.Context) (err error)  {
	var order domain.Order
	err = c.Bind(&order)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": err.Error(),
		})
	}

	// Check and update stock at Product Service
	prodID := order.ProductID
	qty := order.Qty

	//Encode the data
	postBody, _ := json.Marshal(map[string]int{
		"qty": qty,
	})

	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(viper.GetString(`product.url`)+strconv.Itoa(int(prodID)), "application/json", responseBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"err": err.Error(),
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		ctx := c.Request().Context()
		err = oc.OrderService.Store(ctx, &order)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"err": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "order created",
		})
	} else {
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"err": err.Error(),
			})
		}
		sb := string(body)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": sb,
		})
	}

}