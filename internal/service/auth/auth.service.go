package auth

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/internal/repo/gorm"
	migration "github.com/putranurf/be-tpd/internal/repo/migration/auth"
	"github.com/putranurf/be-tpd/pkg/helpers"
)

//Hash Password
func CreateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)
	return c.JSON(http.StatusOK, hash)
}

//Login
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

//Reset Password
func CreateReset(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	newPassword := c.FormValue("newpassword")
	confirmPassword := c.FormValue("confirmpassword")
	if newPassword != confirmPassword {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrBadRequest.Code,
			"message": "New Password and Confirm Password not the same",
		})
	}
	hash, _ := helpers.HashPassword(newPassword)
	arrobj := &migration.Password{
		Id:          id,
		Newpassword: &hash,
	}
	_, err := gorm.CreateReset(arrobj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrBadRequest.Code,
			"message": echo.ErrBadGateway.Message,
		})
	}
	result, _ := gorm.ListUser()
	return c.JSON(http.StatusOK, result)
}

//Register User
func CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	hash, _ := helpers.HashPassword(password)
	_, err := gorm.CreateUser(username, hash, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusConflict,
			"message": "Error conflict (username) request payload",
		})
	}
	result, _ := gorm.ListUser()
	return c.JSON(http.StatusOK, result)
}
