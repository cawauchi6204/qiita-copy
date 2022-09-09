package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/test_db?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var DB = dbConnect()
