package http

import (
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/internal/service/user"
	middlewares "github.com/putranurf/be-tpd/pkg/middleware"
)

func InitUser(group *echo.Group) {
	// User API
	group.GET("/list-user", user.ListUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	group.PUT("/update-user/:id", user.UpdateUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	group.PUT("/delete-user/:id", user.DeleteUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

}
