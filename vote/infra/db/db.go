package db

import (
	"gomq/vote/entities"
	"log"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

func ConnectDB() *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	db, err = gorm.Open("vote.db", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&entities.Vote{})

	return db
}
