package controllers

import (
	"golang_poc/config"
	"golang_poc/models"
	"net/http"
	// "encoding/json"

	"github.com/gin-gonic/gin"
)
// Create Multiple Employees
func CreateEmployees(c *gin.Context) {
	var employees []models.Employee // Ensure this is a slice

	// Bind JSON array to employees slice
	if err := c.ShouldBindJSON(&employees); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if any employee with the same email already exists
	for _, emp := range employees {
		var existingEmployee models.Employee
		if err := config.DB.Where("email = ?", emp.Email).First(&existingEmployee).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Employee with email " + emp.Email + " already exists"})
			return
		}
	}

	// Bulk insert employees
	result := config.DB.Create(&employees)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employees added successfully", "employees": employees})
}


// Get All Employees
func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	config.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

// Get Employee by ID
func GetEmployeeByID(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")
	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// Update Employee
// Update Employee
func UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	// Find existing employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Create a temporary struct to hold update values
	var updateData struct {
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Position string  `json:"position"`
		Age      int     `json:"age"`
		Salary   float64 `json:"salary"`
	}

	// Bind JSON to updateData struct
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only non-empty fields
	if updateData.Name != "" {
		employee.Name = updateData.Name
	}
	if updateData.Email != "" {
		employee.Email = updateData.Email
	}
	if updateData.Position != "" {
		employee.Position = updateData.Position
	}
	if updateData.Age != 0 {
		employee.Age = updateData.Age
	}
	if updateData.Salary != 0 {
		employee.Salary = updateData.Salary
	}

	// Save updated employee
	config.DB.Save(&employee)
	c.JSON(http.StatusOK, employee)
}


// Delete Employee
func DeleteEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")
	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	config.DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
