package db

import (
	"gomq/user/entities"
	"log"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

func ConnectDB() *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	db, err = gorm.Open("user.db", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&entities.User{})

	return db
}
