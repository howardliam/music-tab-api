package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {
	return c.String(http.StatusOK, "Music Tablature API")
}
