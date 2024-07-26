package database

import (
	"CRUD-Golang/config"
	"CRUD-Golang/models"
)

func InitDatabase() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Doctor{}, &models.Patient{})
}
