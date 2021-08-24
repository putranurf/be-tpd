package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/db"
	internalhttp "github.com/putranurf/be-tpd/internal/http"
)

func main() {
	//Env Init
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//DB Init
	db.Init()

	// Routes Init
	e := echo.New()
	apiGroup := e.Group("/api")
	internalhttp.InitAuth(apiGroup)
	internalhttp.InitTest(apiGroup)
	internalhttp.InitUser(apiGroup)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
