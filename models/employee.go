package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Age      int     `json:"age"`
	Salary   float64 `json:"salary"`
	PhoneNumber string  `json:"phone_number"`
}
