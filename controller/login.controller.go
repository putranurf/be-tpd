package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/helpers"
	"github.com/putranurf/be-tpd/model/gorm"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := gorm.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrUnauthorized.Code,
			"message": echo.ErrUnauthorized.Message,
		})
	}
	return c.JSON(http.StatusOK, res)
}
