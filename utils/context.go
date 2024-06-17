package utils

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TabContext struct {
	echo.Context
	DB *gorm.DB
}
