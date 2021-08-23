package http

import (
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/internal/service/auth"
	middlewares "github.com/putranurf/be-tpd/pkg/middleware"
)

func InitAuth() {
	e := echo.New()

	// Auth API
	e.GET("/create-hash/:password", auth.CreateHashPassword, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.POST("/auth", auth.GetLogin, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.POST("/reset/:id", auth.CreateReset, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.POST("/create-user", auth.CreateUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

}
