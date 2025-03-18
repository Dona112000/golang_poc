package controllers

import (
	"bytes"
	// "encoding/json"
	"golang_poc/config"
	"golang_poc/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup test database
func setupTestDB() {
	var err error
	config.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	// Auto-migrate tables
	config.DB.AutoMigrate(&models.Employee{})
}

// Helper function to create a test router
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/employees", CreateEmployees)
	r.GET("/employees", GetEmployees)
	r.GET("/employees/:id", GetEmployeeByID)
	r.PUT("/employees/:id", UpdateEmployee)
	r.DELETE("/employees/:id", DeleteEmployee)
	return r
}

// Run setup before tests
func TestMain(m *testing.M) {
	setupTestDB()
	code := m.Run()
	os.Exit(code)
}

// Test Create Employee API
func TestCreateEmployees(t *testing.T) {
	r := setupRouter()

	employeeData := `[{"name": "John Doe", "email": "john@example.com", "position": "Developer", "age": 30, "salary": 60000}]`
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer([]byte(employeeData)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// Test Get All Employees API
func TestGetEmployees(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/employees", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// Test Get Employee by ID
func TestGetEmployeeByID(t *testing.T) {
	r := setupRouter()

	// Create employee first
	employee := models.Employee{Name: "Alice", Email: "alice@example.com", Position: "Manager", Age: 40, Salary: 75000}
	config.DB.Create(&employee)

	req, _ := http.NewRequest("GET", "/employees/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// Test Update Employee API
func TestUpdateEmployee(t *testing.T) {
	r := setupRouter()

	// Create employee first
	employee := models.Employee{Name: "Bob", Email: "bob@example.com", Position: "Analyst", Age: 35, Salary: 50000}
	config.DB.Create(&employee)

	updateData := `{"name": "Bob Updated", "position": "Senior Analyst", "salary": 55000}`
	req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer([]byte(updateData)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// Test Delete Employee API
func TestDeleteEmployee(t *testing.T) {
	r := setupRouter()

	// Create employee first
	employee := models.Employee{Name: "Eve", Email: "eve@example.com", Position: "HR", Age: 28, Salary: 40000}
	config.DB.Create(&employee)

	req, _ := http.NewRequest("DELETE", "/employees/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
