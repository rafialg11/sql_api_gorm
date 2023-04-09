package database

import (
	"fmt"
	"log"
	"sql-api-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	port     = "5432"
	dbname   = "book_gorm"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to DB :", err)
	}

	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
