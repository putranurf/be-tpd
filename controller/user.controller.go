package controller

import (
	"github.com/putranurf/be-tpd/model/gorm"

	"net/http"

	"github.com/labstack/echo"
)

// GetUsers
func GetUsers(c echo.Context) error {
	result, err := gorm.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
