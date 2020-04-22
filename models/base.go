package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	// Build connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	//  &Contact{}
	db.Debug().AutoMigrate(&Account{}) // Database Migration. Debug is logger
}

func GetDB() *gorm.DB {
	return db
}
