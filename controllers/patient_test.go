package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	// Set up Gin router and routes for testing

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/patient/", CreatePatient)

	// Define the request payload
	payload := []byte(`{
		"name": "Yuji Itadori The true MC",
		"contact_no": "0987654321",
		"address": "Sendai, Japan",
		"doctor_id": "724eec76-4cee-11ef-8de2-6b492c02869b"
	}`)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/patient/", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder to record the response
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	// Since the ID is generated randomly, we'll need to match the rest of the response
	var actualResponse map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := map[string]interface{}{
		"name":       "Yuji Itadori The true MC",
		"contact_no": "0987654321",
		"address":    "Sendai, Japan",
		"doctor_id":  "724eec76-4cee-11ef-8de2-6b492c02869b",
	}

	// Ensure ID and timestamps are present, and other fields match
	assert.Contains(t, actualResponse, "id")
	assert.Contains(t, actualResponse, "created_at")
	assert.Contains(t, actualResponse, "updated_at")
	assert.Equal(t, expectedResponse["name"], actualResponse["name"])
	assert.Equal(t, expectedResponse["contact_no"], actualResponse["contact_no"])
	assert.Equal(t, expectedResponse["address"], actualResponse["address"])
	assert.Equal(t, expectedResponse["doctor_id"], actualResponse["doctor_id"])
}

func TestGetPatientByID(t *testing.T) {
	// Set up Gin router and routes for testing
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/patient/:id", GetPatientByID)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/patient/724f87bc-4cee-11ef-8de2-6b492c02869b", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP recorder to record the response
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	expected := `{
		"id":"724f87bc-4cee-11ef-8de2-6b492c02869b",
		"created_at":"2024-07-28T20:04:07+05:30",
		"updated_at":"2024-07-28T20:04:07+05:30",
		"name":"Maki Zenin",
		"contact_no":"2345678901",
		"address":"Akihabara, Tokyo",
		"doctor_id":"724eec76-4cee-11ef-8de2-6b492c02869b",
		"DeletedAt":null
	}`
	assert.JSONEq(t, expected, rr.Body.String())
}

// Similar tests will be written for the other functions like UpdatePatient, DeletePatient, etc.
