package routes

import (
	"golang_poc/controllers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine) {
	router.POST("/employees", controllers.CreateEmployees)
	router.GET("/employees", controllers.GetEmployees)
	router.GET("/employees/:id", controllers.GetEmployeeByID)
	router.PUT("/employees/:id", controllers.UpdateEmployee)
	router.DELETE("/employees/:id", controllers.DeleteEmployee)
}
