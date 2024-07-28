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
	router.DELETE("/doctor/:id", controllers.DeleteDoctor)
	router.GET("/searchDoctorByName", controllers.SearchDoctorByName)

	// Patient Routes
	router.POST("/patient/", controllers.CreatePatient)
	router.GET("/patient/:id", controllers.GetPatientByID)
	router.PATCH("/patient/:id", controllers.UpdatePatient)
	router.GET("/fetchPatientByDoctorId/:doctor_id", controllers.GetPatientsByDoctorID)
	router.DELETE("/patient/:id", controllers.DeletePatient)
	router.GET("/searchPatientByName", controllers.SearchPatientByName)
	return router
}
