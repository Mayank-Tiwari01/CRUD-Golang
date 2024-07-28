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

	if err := config.DB.First(&patient, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found!"})
		return
	}

	var input struct {
		Name      string `json:"name"`
		ContactNo string `json:"contact_no"`
		Address   string `json:"address"`
		DoctorID  string `json:"doctor_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient.Name = input.Name
	patient.ContactNo = input.ContactNo
	patient.Address = input.Address
	patient.DoctorID = input.DoctorID
	patient.UpdatedAt = time.Now()

	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
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
func GetPatientsByDoctorIDOrName(c *gin.Context) {
	var patients []models.Patient
	doctorID := c.Query("doctor_id")
	doctorName := c.Query("doctor_name")

	if doctorID != "" {
		if err := config.DB.Where("doctor_id = ?", doctorID).Find(&patients).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No patients found for the given doctor ID"})
			return
		}
	} else if doctorName != "" {
		var doctor models.Doctor
		if err := config.DB.Where("name = ?", doctorName).First(&doctor).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found!"})
			return
		}
		if err := config.DB.Where("doctor_id = ?", doctor.ID).Find(&patients).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No patients found for the given doctor name"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a doctor ID or name"})
		return
	}

	if len(patients) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No patients found"})
		return
	}

	c.JSON(http.StatusOK, patients)
}
