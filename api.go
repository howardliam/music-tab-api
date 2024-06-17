package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/howardliam/music-tab-api/config"
	"github.com/howardliam/music-tab-api/database"
	"github.com/howardliam/music-tab-api/routes"
	"github.com/howardliam/music-tab-api/security"
	"github.com/howardliam/music-tab-api/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	// Load configuration
	conf := config.LoadConfig()

	// Connect to database and migrate
	db := database.NewDatabase(conf.Postgres)
	database.Migrate(db)

	// Connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Unable to establish connection pool: %v", err)
	}
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Initialise the web server
	e := echo.New()

	// Inject database into context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tc := &utils.TabContext{Context: c, DB: db}
			return next(tc)
		}
	})

	// Logger
	logger := zerolog.New(os.Stdout)
	loggerConfig := middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().Str("URI", v.URI).Int("status", v.Status).Str("IP", v.RemoteIP).Msg("request")
			return nil
		},
	}
	e.Use(middleware.RequestLoggerWithConfig(loggerConfig))
	// e.Use(middleware.Logger())

	// Error recovery
	e.Use(middleware.Recover())

	// Error handler
	e.HTTPErrorHandler = routes.ErrorHandler

	// Icon
	e.File("/favicon.ico", "static/images/favicon.ico")

	// Routes
	e.GET("/", routes.Root)

	auth := e.Group("/login")
	auth.POST("", routes.Login)

	/* Secured routes start */
	jwtConf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(security.JWTClaims)
		},
		SigningKey: []byte("SuperSecret"),
	}

	// Bands routes
	bands := e.Group("/bands")
	bands.Use(echojwt.WithConfig(jwtConf))
	bands.GET("", routes.GetAllBands)
	bands.GET("/:id", routes.GetBandById)

	// Albums routes
	albums := e.Group("/albums")
	albums.Use(echojwt.WithConfig(jwtConf))

	// Songs routes
	songs := e.Group("/songs")
	songs.Use(echojwt.WithConfig(jwtConf))

	// Tabs routes
	tabs := e.Group("/tabs")
	tabs.Use(echojwt.WithConfig(jwtConf))
	/* Secured routes end */

	// Start web server
	// if err := e.StartTLS(utils.GenerateAddress(conf.Server), "cert", "key"); err != http.ErrServerClosed {
	if err := e.Start(utils.GenerateAddress(conf.Server)); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
