package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func InitDb() *gorm.DB {
	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(driver, dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
