package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("tournament.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return database
}
