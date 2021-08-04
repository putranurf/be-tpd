package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/controller"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/users", controller.GetUsers)
	return e
}
