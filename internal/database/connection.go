package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectToDB establishes the connection to the DB
// the connection string is still hardcoded by now
func ConnectToDB() {
	dsn := "host=localhost user=postgres password=265ameth dbname=longodb port=5432 sslmode=disable"
	dbg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error establising coneection to DB")
	}
	DB = dbg

	fmt.Println("Connection to DB establised")

}
