package user

import (
	"net/http"
)

// GetUsers
// func GetUsers(c echo.Context) error {
// 	// return func(c echo.Context) error {
// 	// userCookie, _ := c.Cookie("user")
// 	result, err := gorm.GetUsers()

// 	for key, values := range c.Request().Header {
// 		fmt.Println(key)
// 		for _, value := range values {
// 			fmt.Println(value)
// 		}
// 	}

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"status":  echo.ErrUnauthorized.Code,
// 			"message": echo.ErrUnauthorized.Message,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, result)
// 	// return c.String(http.StatusOK, "Hello World")
// 	// }
// // }

func GetUsers(w http.ResponseWriter, r *http.Request) {

}
