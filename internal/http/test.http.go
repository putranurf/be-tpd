package http

import (
	"net/http"

	"github.com/labstack/echo"
	controller "github.com/putranurf/be-tpd/internal/service"
	middlewares "github.com/putranurf/be-tpd/pkg/middleware"
)

func InitTest(group *echo.Group) {
	//Testing API
	group.GET("/index", HandlerIndex, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	group.GET("/test-struct", controller.TestStructValidation, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	group.GET("/test-var", controller.TestVarValidation, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
}

func HandlerIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Index API")
}
