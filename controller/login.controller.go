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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "berhasil login")
}
