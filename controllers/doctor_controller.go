package controllers

import (
	"net/http"
	"time"

	"CRUD-Golang/config"
	"CRUD-Golang/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDoctor(c *gin.Context) {
	var input models.Doctor
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

func GetDoctorByID(c *gin.Context) {
	var doctor models.Doctor
	if err := config.DB.Where("id = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found!"})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func UpdateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := config.DB.Where("id = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found!"})
		return
	}

	var input struct {
		ContactNo string `json:"contact_no"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&doctor).Updates(models.Doctor{ContactNo: input.ContactNo, UpdatedAt: time.Now()})

	c.JSON(http.StatusOK, doctor)
}
