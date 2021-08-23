package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/db"
	internalhttp "github.com/putranurf/be-tpd/internal/http"
	"github.com/putranurf/be-tpd/internal/service/token"
	middlewares "github.com/putranurf/be-tpd/pkg/middleware"
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

	// Routes Init
	e.POST("/create-token", token.CreateToken, echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	internalhttp.InitAuth()
	internalhttp.InitTest()
	internalhttp.InitUser()

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
