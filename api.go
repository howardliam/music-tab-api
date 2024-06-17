package main

import (
	"log"
	"strconv"
	"time"

	"github.com/brpaz/echozap"
	"github.com/howardliam/music-tab-api/config"
	"github.com/howardliam/music-tab-api/database"
	"github.com/howardliam/music-tab-api/routes"
	"github.com/howardliam/music-tab-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
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

	// Set up middleware
	zapLogger, _ := zap.NewProduction()
	e.Use(echozap.ZapLogger(zapLogger))
	e.Use(middleware.Recover())

	// Error handler
	e.HTTPErrorHandler = routes.ErrorHandler

	// Routes
	e.GET("/", routes.Root)
	e.GET("/bands", routes.GetAllBands)
	e.GET("/bands/:id", routes.GetBandById)

	// Start web server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(conf.Server.Port))))
}
