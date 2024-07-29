package database

import (
	"CRUD-Golang/config"
	"CRUD-Golang/models"
	"testing"
)

func TestSetup(t *testing.T) {
	// the test database
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Doctor{}, &models.Patient{})
}
