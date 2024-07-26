package routers

import (
	"CRUD-Golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Doctor Routes
	router.POST("/doctor/", controllers.CreateDoctor)
	router.GET("/doctor/:id", controllers.GetDoctorByID)
	router.PATCH("/doctor/:id", controllers.UpdateDoctor)

	// Patient Routes
	router.POST("/patient/", controllers.CreatePatient)
	router.GET("/patient/:id", controllers.GetPatientByID)
	router.PATCH("/patient/:id", controllers.UpdatePatient)
	router.GET("/fetchPatientByDoctorId/:doctor_id", controllers.GetPatientsByDoctorID)

	return router
}
