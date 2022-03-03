package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Index get nothing
func Index(c echo.Context) error {
	routes, err := json.MarshalIndent(c.Echo().Routes(), "", " ")
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(routes))
}
