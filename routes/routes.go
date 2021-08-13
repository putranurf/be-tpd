package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/controller"
	"github.com/putranurf/be-tpd/controller/auth"
	"github.com/putranurf/be-tpd/controller/token"
	"github.com/putranurf/be-tpd/controller/user"
	"github.com/putranurf/be-tpd/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	//Token
	e.POST("/fetch-token", token.FetchToken)

	//Auth
	e.POST("/login", auth.CheckLogin, middleware.IsAuthenticated)
	e.GET("/generate-hash/:password", auth.GenerateHashPassword)

	//CRUD
	e.GET("/users", user.GetUsers, middleware.AuthToken())
	e.GET("/test-struct", controller.TestStructValidation)
	e.GET("/test-var", controller.TestVarValidation)

	return e
}
