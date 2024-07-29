package controllers

import (
	"CRUD-Golang/config"
	"CRUD-Golang/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/doctor/", CreateDoctor)
	return router
}

func setupTestDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/test_mayankDB?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	config.DB = database
	config.DB.AutoMigrate(&models.Doctor{}, &models.Patient{})
}

func TestCreateDoctor(t *testing.T) {
	setupTestDatabase()
	defer tearDownTestDatabase()

	router := setupRouter()

	input := models.Doctor{
		Name:      "Dr. Satoru Gojo",
		ContactNo: "1234567890",
		Address:   "Tokyo, Japan",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/doctor/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Doctor
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, input.Name, response.Name)
	assert.Equal(t, input.ContactNo, response.ContactNo)
	assert.Equal(t, input.Address, response.Address)
}

func tearDownTestDatabase() {
	config.DB.Exec("DROP TABLE IF EXISTS doctors")
	config.DB.Exec("DROP TABLE IF EXISTS patients")
}
