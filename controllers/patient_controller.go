package controllers

import (
	"net/http"
	"time"

	"CRUD-Golang/config"
	"CRUD-Golang/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePatient(c *gin.Context) {
	var input models.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = uuid.New().String()[:5]
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	config.DB.Create(&input)

	c.JSON(http.StatusOK, input)
}

func GetPatientByID(c *gin.Context) {
	var patient models.Patient
	if err := config.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found!"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func UpdatePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	// Find the patient in the database using the ID
	if err := config.DB.First(&patient, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found!"})
		return
	}

	// Define an anonymous struct to hold the incoming JSON input
	var input struct {
		Name      *string `json:"name"`
		ContactNo *string `json:"contact_no"`
		Address   *string `json:"address"`
		DoctorID  *string `json:"doctor_id"`
	}

	// Bind the JSON input to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the patient fields only if the input is not nil
	if input.Name != nil {
		patient.Name = *input.Name
	}
	if input.ContactNo != nil {
		patient.ContactNo = *input.ContactNo
	}
	if input.Address != nil {
		patient.Address = *input.Address
	}
	if input.DoctorID != nil {
		patient.DoctorID = *input.DoctorID
	}
	patient.UpdatedAt = time.Now() // Update the UpdatedAt timestamp

	// Save the updated patient to the database
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient) // Return the updated patient as JSON
}

func DeletePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found!"})
		return
	}

	config.DB.Delete(&patient)

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

func GetPatientsByDoctorID(c *gin.Context) {
	var patients []models.Patient
	if err := config.DB.Where("doctor_id = ?", c.Param("doctor_id")).Find(&patients).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patients not found!"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func SearchPatientByName(c *gin.Context) {
	name := c.Query("name")
	var patients []models.Patient
	if err := config.DB.Where("name LIKE ?", "%"+name+"%").Find(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}
