package db

import (
	"fmt"
	"os"

	"github.com/cadenkoj/vera/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() (err error) {
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	port := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database.AutoMigrate(&model.Profile{})
	return err
}

func GetDB() *gorm.DB {
	return database
}
