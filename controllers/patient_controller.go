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
	if err := config.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found!"})
		return
	}

	var input struct {
		ContactNo string `json:"contact_no"`
		Address   string `json:"address"`
		DoctorID  string `json:"doctor_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&patient).Updates(models.Patient{ContactNo: input.ContactNo, Address: input.Address, DoctorID: input.DoctorID, UpdatedAt: time.Now()})

	c.JSON(http.StatusOK, patient)
}

func GetPatientsByDoctorID(c *gin.Context) {
	var patients []models.Patient
	if err := config.DB.Where("doctor_id = ?", c.Param("doctor_id")).Find(&patients).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patients not found!"})
		return
	}

	c.JSON(http.StatusOK, patients)
}
