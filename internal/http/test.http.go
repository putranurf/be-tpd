package http

import (
	"net/http"

	"github.com/labstack/echo"
	controller "github.com/putranurf/be-tpd/internal/service"
	middlewares "github.com/putranurf/be-tpd/pkg/middleware"
)

func InitTest() {
	e := echo.New()

	//Testing API
	e.GET("/index", HandlerIndex, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.GET("/test-struct", controller.TestStructValidation, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.GET("/test-var", controller.TestVarValidation, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
}

func HandlerIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Index API")
}
