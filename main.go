package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/putranurf/be-tpd/db"
	"github.com/putranurf/be-tpd/routes"
)

func main() {
	//Routes Init

	//Env Init
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//DB Init
	db.Init()
	routes.Init()
	// Start server
	// e.Logger.Fatal(e.Start(":1234"))
}
