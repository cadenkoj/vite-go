package main

import (
	"log"

	"github.com/cadenkoj/vera/backend/db"
	"github.com/cadenkoj/vera/backend/router"
	"github.com/cadenkoj/vera/backend/web"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := db.Connect(); err != nil {
		log.Fatal("Error connecting to database")
	}

	router := router.New()
	web.RegisterHandlers(router)

	router.Start(":8080")
}
