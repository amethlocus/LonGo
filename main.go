package main

import (
	"LonGo/internal/database"
	"LonGo/internal/routing"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectToDB()
	// Register all routes
	routing.RegisterRoutes()

}
