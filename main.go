package main

import (
	"golang_poc/config"
	"golang_poc/models"
	"golang_poc/routes"
	// "golang_poc/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDB()

	// Migrate the database
	config.DB.AutoMigrate(&models.Employee{})

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.EmployeeRoutes(router)

	// Start server
	router.Run(":8081")
}
