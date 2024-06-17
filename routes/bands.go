package routes

import (
	"net/http"

	"github.com/howardliam/music-tab-api/models"
	"github.com/howardliam/music-tab-api/utils"
	"github.com/labstack/echo/v4"
)

func GetAllBands(c echo.Context) error {
	tc := c.(*utils.TabContext)

	var result []models.Band
	tc.DB.Model(&models.Band{}).Find(&result)

	return c.JSON(http.StatusOK, result)
}

func GetBandById(c echo.Context) error {
	tc := c.(*utils.TabContext)

	param := tc.Param("id")
	var result models.Band
	err := tc.DB.Model(&models.Band{}).First(&result, param).Error
	if err != nil {
		res := utils.Response{
			StatusCode: http.StatusNotFound,
			Data:       nil,
			Error: utils.Error{
				Message: "Not found",
			},
		}
		return c.JSON(http.StatusNotFound, res)
	}

	return c.JSON(http.StatusOK, result)
}
