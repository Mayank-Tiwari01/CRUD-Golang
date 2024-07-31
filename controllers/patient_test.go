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

func setup() {
	mockPatients = make(map[string]models.Patient)
}

func TestCreatePatient(t *testing.T) {
	setup()

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/patient", func(c *gin.Context) {
		var input models.Patient
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input.ID = "mock_id_" + input.ID
		input.CreatedAt = time.Now()
		input.UpdatedAt = time.Now()

		MockDBCreatePatient(&input)

		c.JSON(http.StatusOK, input)
	})

	newPatient := models.Patient{
		Name:      "Jane Doe",
		ContactNo: "0987654321",
		Address:   "123 Main St",
		DoctorID:  "doctor_id_123",
	}

	jsonValue, _ := json.Marshal(newPatient)
	req, _ := http.NewRequest("POST", "/patient", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var createdPatient models.Patient
	err := json.Unmarshal(resp.Body.Bytes(), &createdPatient)
	assert.NoError(t, err)
	assert.Equal(t, newPatient.Name, createdPatient.Name)
}

func TestGetPatientByID(t *testing.T) {
	setup()

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/patient/:id", func(c *gin.Context) {
		id := c.Param("id")
		patient, err := MockDBFindPatientByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found!"})
			return
		}
		c.JSON(http.StatusOK, patient)
	})

	mockPatients["mock_id_123"] = models.Patient{
		ID:        "mock_id_123",
		Name:      "Jane Doe",
		ContactNo: "0987654321",
		Address:   "123 Main St",
		DoctorID:  "doctor_id_123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	req, _ := http.NewRequest("GET", "/patient/mock_id_123", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var patient models.Patient
	err := json.Unmarshal(resp.Body.Bytes(), &patient)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", patient.Name)
}
