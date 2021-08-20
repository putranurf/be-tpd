package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/repo/gorm"
)

//Create User
func CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	_, err := gorm.CreateUser(username, password, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrUnauthorized.Code,
			"message": echo.ErrUnauthorized.Message,
		})
	}
	result, _ := gorm.ListUser()
	return c.JSON(http.StatusOK, result)
}

// Get List User
func ListUser(c echo.Context) error {
	result, err := gorm.ListUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrUnauthorized.Code,
			"message": echo.ErrUnauthorized.Message,
		})
	}
	return c.JSON(http.StatusOK, result)
}

// Update User
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	username := c.FormValue("username")
	password := c.FormValue("password")
	_, err := gorm.UpdateUser(id, username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrUnauthorized.Code,
			"message": echo.ErrUnauthorized.Message,
		})
	}
	result, _ := gorm.ListUser()
	return c.JSON(http.StatusOK, result)
}

// Delete User
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := gorm.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrUnauthorized.Code,
			"message": echo.ErrUnauthorized.Message,
		})
	}
	result, _ := gorm.ListUser()
	return c.JSON(http.StatusOK, result)
}
