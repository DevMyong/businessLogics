package api

import (
	"businessLogics/model/price"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetPrices get price from in-memory cache (concurrent-map)
func GetPrices(c echo.Context) error {
	exchange := c.QueryParam("exchange")
	prices := price.GetAll(exchange)

	return c.JSONPretty(http.StatusOK, prices, " ")
}
