package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbPort := os.Getenv("db_port")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	if e != nil {
		fmt.Print(e)
	}

	conn, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, username, dbName, password))

	if err != nil {
		fmt.Print(err)
	}
	db = conn
}

func Postgres() *gorm.DB {
	return db
}
