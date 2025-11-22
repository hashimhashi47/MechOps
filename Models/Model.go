package models

import "gorm.io/gorm"

// user
type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	Block     bool   `gorm:"default:false"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
}
