package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/controller"
	"github.com/putranurf/be-tpd/controller/auth"
	"github.com/putranurf/be-tpd/controller/token"
	"github.com/putranurf/be-tpd/controller/user"
	"github.com/putranurf/be-tpd/db"
	middlewares "github.com/putranurf/be-tpd/middleware"
)

func main() {
	//Routes Init
	e := echo.New()

	//Env Init
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//DB Init
	db.Init()

	//Routes Init
	e.POST("/create-token", token.CreateToken, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

	e.GET("/create-hash/:password", auth.CreateHashPassword, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.POST("/auth", auth.GetLogin, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

	e.GET("/list-user", user.ListUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.PUT("/update-user/:id", user.UpdateUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.POST("/create-user", user.CreateUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.PUT("/delete-user/:id", user.DeleteUser, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

	//Testing API
	e.GET("/index", HandlerIndex, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.GET("/test-struct", controller.TestStructValidation, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.GET("/test-var", controller.TestVarValidation, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

	fmt.Println("Masuk BE")
	// Start server
	e.Logger.Fatal(e.Start(":1234"))

}

func HandlerIndex(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hello Index API")

}
