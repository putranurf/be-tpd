package db

import (
	"log"

	"github.com/putranurf/be-tpd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB a global for GORM DB object
var DB *gorm.DB

//Init to initiate Database Connection
func Init() *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	// DB, err = gorm.Open(config.GetDBType(), conString)
	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB
}

//GetDBInstance return DB object
func GetDBInstance() *gorm.DB {
	return DB
}
