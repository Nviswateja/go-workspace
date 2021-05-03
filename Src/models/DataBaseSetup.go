package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Setup() {
	//dsn := "host=localhost user=postgres password=Password1 dbname=test port=5432 sslmode=disable"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "postgres", "test", "Password1") //Build connection string
	fmt.Println(dbUri)
	database, err := gorm.Open("postgres", dbUri)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Customer{})

	DB = database
}
