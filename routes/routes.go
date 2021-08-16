package routes

import (
	"net/http"

	// "github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"

	// "github.com/labstack/echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/putranurf/be-tpd/controller/token"
	"github.com/putranurf/be-tpd/controller/user"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	// // Token
	e.POST("/fetch-token", token.FetchToken)

	// // Auth
	// // e.POST("/login", auth.CheckLogin, middleware.IsAuthenticated)
	// e.GET("/generate-hash/:password", auth.GenerateHashPassword)

	// // CRUD
	e.GET("/users", user.GetUsers)
	// e.GET("/test-struct", controller.TestStructValidation)
	// e.GET("/test-var", controller.TestVarValidation)

	// adminGroup := e.Group("/v1")
	// config := middleware.JWTConfig{
	// 	// Claims:     &migration.JwtCustomClaims{},
	// 	SigningKey: []byte(os.Getenv("TOKEN_SECRET")),
	// }
	// adminGroup.Use(middleware.JWTWithConfig(config))

	// // // Attach jwt token refresher.
	// // adminGroup.Use(middlewares.TokenRefresherMiddleware)

	// adminGroup.GET("/users", user.GetUsers())

	return e
}
