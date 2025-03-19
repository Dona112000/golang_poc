package controllers

import (
	"golang_poc/config"
	"golang_poc/models"
	"net/http"
	// "encoding/json"

	"github.com/gin-gonic/gin"
)
func CreateEmployees(c *gin.Context) {
	var employees []models.Employee

	// Bind JSON array to employees slice
	if err := c.ShouldBindJSON(&employees); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, emp := range employees {
		var existingEmployee models.Employee
		// Check if an employee with the same email or phone number exists
		if err := config.DB.Where("email = ? OR phone_number = ?", emp.Email, emp.PhoneNumber).First(&existingEmployee).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Employee with email " + emp.Email + " or phone number " + emp.PhoneNumber + " already exists"})
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

func UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var updateData struct {
		Name        string  `json:"name"`
		Email       string  `json:"email"`
		Position    string  `json:"position"`
		Age         int     `json:"age"`
		Salary      float64 `json:"salary"`
		PhoneNumber string  `json:"phone_number"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	if updateData.PhoneNumber != "" {
		employee.PhoneNumber = updateData.PhoneNumber
	}

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
