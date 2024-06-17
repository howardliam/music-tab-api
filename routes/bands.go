package routes

import (
	"net/http"

	"github.com/howardliam/music-tab-api/models"
	"github.com/howardliam/music-tab-api/utils"
	"github.com/labstack/echo/v4"
)

func GetAllBands(c echo.Context) error {
	tc := c.(*utils.TabContext)

	var bands []models.Band
	tc.DB.Model(&models.Band{}).Find(&bands)

	res := utils.Response{
		StatusCode: http.StatusOK,
		Data:       bands,
		Error:      nil,
	}

	return c.JSON(http.StatusOK, res)
}

func GetBandById(c echo.Context) error {
	tc := c.(*utils.TabContext)

	param := tc.Param("id")
	var band models.Band
	err := tc.DB.Model(&models.Band{}).First(&band, param).Error
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

	res := utils.Response{
		StatusCode: http.StatusOK,
		Data:       band,
		Error:      nil,
	}

	return c.JSON(http.StatusOK, res)
}
