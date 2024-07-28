package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
We use a pointer here so that we can reference the same `gorm.DB` object efficiently across the entire application.
This avoids the overhead of copying the database connection object.
`DB` is a global variable, so other files and packages within the same Go module can access and use this database connection for performing database operations.
*/
var DB *gorm.DB

func ConnectDatabase() {
	//using default credentials
	//dsn is short for Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/mayankDB?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//this ensures that the application does not continue running without a database connection.
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB = database
}
