package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/controller"
	"github.com/putranurf/be-tpd/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/users", controller.GetUsers, middleware.IsAuthenticated)

	e.GET("/generate-hash/:password", controller.GenerateHashPassword)
	e.POST("/login", controller.CheckLogin)

	return e
}
