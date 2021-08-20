package auth

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/helpers"
	"github.com/putranurf/be-tpd/repo/gorm"
)

func CreateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)
	return c.JSON(http.StatusOK, hash)
}

func GetLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := gorm.GetLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrBadRequest.Code,
			"message": echo.ErrBadGateway.Message,
		})
	}

	return c.JSON(http.StatusOK, res)
}

// func CreateForgot(c echo.Context) error {
// 	email := c.FormValue("email")
// 	res, err := gorm.CreateForgot(email)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"status":  echo.ErrBadRequest.Code,
// 			"message": echo.ErrBadGateway.Message,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, res)
// }
