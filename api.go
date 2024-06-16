package main

import (
	"net/http"
	"strconv"

	"github.com/howardliam/music-tab-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	conf := config.LoadConfig()

	// Initialise the web server
	e := echo.New()

	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", root)

	// Start web server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(conf.Server.Port))))
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Music Tablature API")
}
