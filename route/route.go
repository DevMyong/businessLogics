package route

import (
	"businessLogics/controller/api"
	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Index)
	e.GET("/price", api.GetPrices)

	return e
}
