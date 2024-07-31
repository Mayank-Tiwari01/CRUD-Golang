package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"CRUD-Golang/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateDoctor(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/doctor", func(c *gin.Context) {
		var input models.Doctor
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input.ID = "mock_id_" + input.ID
		input.CreatedAt = time.Now()
		input.UpdatedAt = time.Now()

		MockDBCreateDoctor(&input)

		c.JSON(http.StatusOK, input)
	})

	// Test data
	newDoctor := models.Doctor{
		Name:      "John Doe",
		ContactNo: "1234567890",
	}

	jsonValue, _ := json.Marshal(newDoctor)
	req, _ := http.NewRequest("POST", "/doctor", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Perform the test
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var createdDoctor models.Doctor
	err := json.Unmarshal(resp.Body.Bytes(), &createdDoctor)
	assert.NoError(t, err)
	assert.Equal(t, newDoctor.Name, createdDoctor.Name)
}

func TestGetDoctorByID(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/doctor/:id", func(c *gin.Context) {
		id := c.Param("id")
		doctor, err := MockDBFindDoctorByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found!"})
			return
		}
		c.JSON(http.StatusOK, doctor)
	})

	// Add a doctor to the mock database
	mockDoctors["mock_id_123"] = models.Doctor{
		ID:        "mock_id_123",
		Name:      "John Doe",
		ContactNo: "1234567890",
	}

	// Perform the test
	req, _ := http.NewRequest("GET", "/doctor/mock_id_123", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var doctor models.Doctor
	err := json.Unmarshal(resp.Body.Bytes(), &doctor)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", doctor.Name)
}

func TestUpdateDoctor(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.PATCH("/doctor/:id", func(c *gin.Context) {
		id := c.Param("id")
		doctor, err := MockDBFindDoctorByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found!"})
			return
		}

		var input struct {
			Name      string `json:"name"`
			ContactNo string `json:"contact_no"`
			Address   string `json:"address"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		doctor.Name = input.Name
		doctor.ContactNo = input.ContactNo
		doctor.Address = input.Address
		doctor.UpdatedAt = time.Now()

		MockDBSaveDoctor(&doctor)
		c.JSON(http.StatusOK, doctor)
	})

	// Add a doctor to the mock database
	mockDoctors["mock_id_123"] = models.Doctor{
		ID:        "mock_id_123",
		Name:      "John Doe",
		ContactNo: "1234567890",
	}

	// Test data
	updateData := map[string]string{
		"name":       "John Updated",
		"contact_no": "0987654321",
		"address":    "123 New Address",
	}
	jsonValue, _ := json.Marshal(updateData)
	req, _ := http.NewRequest("PATCH", "/doctor/mock_id_123", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Perform the test
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var updatedDoctor models.Doctor
	err := json.Unmarshal(resp.Body.Bytes(), &updatedDoctor)
	assert.NoError(t, err)
	assert.Equal(t, "John Updated", updatedDoctor.Name)
	assert.Equal(t, "0987654321", updatedDoctor.ContactNo)
	assert.Equal(t, "123 New Address", updatedDoctor.Address)
}

func TestDeleteDoctor(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.DELETE("/doctor/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := MockDBDeleteDoctor(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
	})

	// Add a doctor to the mock database
	mockDoctors["mock_id_123"] = models.Doctor{
		ID: "mock_id_123",
	}

	// Perform the test
	req, _ := http.NewRequest("DELETE", "/doctor/mock_id_123", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.NotContains(t, mockDoctors, "mock_id_123")
}

func TestSearchDoctorByName(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/doctor/search", func(c *gin.Context) {
		name := c.Query("name")
		doctors := MockDBSearchDoctorsByName(name)
		c.JSON(http.StatusOK, doctors)
	})

	// Add doctors to the mock database
	mockDoctors["mock_id_123"] = models.Doctor{
		ID:   "mock_id_123",
		Name: "John Doe",
	}
	mockDoctors["mock_id_456"] = models.Doctor{
		ID:   "mock_id_456",
		Name: "Jane Doe",
	}

	// Perform the test
	req, _ := http.NewRequest("GET", "/doctor/search?name=John Doe", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var doctors []models.Doctor
	err := json.Unmarshal(resp.Body.Bytes(), &doctors)
	assert.NoError(t, err)
	assert.Len(t, doctors, 1)
	assert.Equal(t, "John Doe", doctors[0].Name)
}
