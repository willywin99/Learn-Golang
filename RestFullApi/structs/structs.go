package structs

import "github.com/jinzhu/gorm"

// Person represent the model for an person
type Person struct {
	gorm.Model
	First_Name string `json:"first_name" example:"first"`
	Last_Name  string `json:"last_name" example:"user"`
}
