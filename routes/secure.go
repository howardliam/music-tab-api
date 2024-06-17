package routes

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/howardliam/music-tab-api/security"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	claims := &security.JWTClaims{
		Name: "Harold",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8760)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("SuperSecret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
