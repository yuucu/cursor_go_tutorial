package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler handles the /hello endpoint
func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
