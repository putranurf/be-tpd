package http

import (
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/internal/service/user"
	middlewares "github.com/putranurf/be-tpd/pkg/middleware"
)

func InitUser() {

	e := echo.New()

	// User API
	e.GET("/list-user", user.ListUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.PUT("/update-user/:id", user.UpdateUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.PUT("/delete-user/:id", user.DeleteUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
}
