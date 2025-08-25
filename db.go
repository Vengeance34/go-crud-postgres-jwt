package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	// ⚠️ Adjust your DSN here
	dsn := "host=localhost user=crud_user password=crud_password dbname=crud port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&User{})
}
