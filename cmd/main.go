package main

import (
	"shortener/internal/database"
	"shortener/internal/handlers"
)

func main() {
	db := database.ConnectDB()
	router := handlers.CreateRouter(db)

	router.Run(":8080")
}
