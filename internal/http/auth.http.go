package http

import (
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/internal/service/auth"
	"github.com/putranurf/be-tpd/internal/service/token"
	middlewares "github.com/putranurf/be-tpd/pkg/middleware"
)

func InitAuth(group *echo.Group) {
	group.POST("/create-token", token.CreateToken, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	// Auth API
	group.GET("/create-hash/:password", auth.CreateHashPassword, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	group.POST("/auth", auth.GetLogin, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	group.POST("/reset/:id", auth.CreateReset, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	group.POST("/create-user", auth.CreateUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
}
