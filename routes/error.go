package routes

import (
	"github.com/howardliam/music-tab-api/utils"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	httpErr := err.(*echo.HTTPError)

	res := utils.Response{
		StatusCode: httpErr.Code,
		Data:       nil,
		Error: utils.Error{
			Message: httpErr.Message.(string),
		},
	}

	c.JSON(httpErr.Code, res)
}
