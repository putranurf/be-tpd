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

	token := c.Response().Header().Get("token")
	restoken, err := gorm.GetToken()

	// r := map[string]interface{}{}
	// r["token"] = restoken.Data
	// e, err := json.Marshal(r)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(e))
	fmt.Println(token)
	fmt.Println(restoken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
