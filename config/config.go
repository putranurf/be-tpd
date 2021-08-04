//GetPostgresConnectionString for GORM
package config

import (
	"fmt"
	"os"
)

//Get DB Type
func GetDBType() string {
	return os.Getenv("DB_TYPE")
}

//Get Connection DB Server
func GetPostgresConnectionString() string {
	db := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))
	return db
}
