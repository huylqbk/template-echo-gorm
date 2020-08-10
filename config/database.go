package config

import (
	"log"
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Database() *gorm.DB {
	driver := os.Getenv("DB_DRIVER")
	database := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	
	var err error
	DB, err = gorm.Open(driver, database)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Database Connected")

	return DB
}