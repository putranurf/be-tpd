package user

import (
	"fmt"

	"github.com/putranurf/be-tpd/model/gorm"

	"net/http"

	"github.com/labstack/echo"
)

// GetUsers
func GetUsers(c echo.Context) error {
	result, err := gorm.GetUsers()

	for key, values := range c.Request().Header {
		fmt.Println(key)
		for _, value := range values {
			fmt.Println(value)
		}
	}
	// restoken, err := gorm.GetToken()

	// token := c.Response().Header().Get("token")

	// fmt.Println(token)
	// fmt.Println(restoken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrUnauthorized.Code,
			"message": echo.ErrUnauthorized.Message,
		})
	}
	return c.JSON(http.StatusOK, result)
}
