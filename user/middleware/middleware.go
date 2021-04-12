package middleware

import "github.com/labstack/echo/v4"

type GoMiddleware struct {

}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

func (mid *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc  {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}