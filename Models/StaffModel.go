package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	Email        string    `json:"email" gorm:"unique"`
	Role         string    `json:"role"`
	Address      []Address `json:"address" gorm:"foreignKey:StaffID;constraint:OnDelete:CASCADE"`
	IdentityCard string    `json:"cardnumber"`
	Block        bool      `gorm:"default:false"`
	Password     string    `json:"password"`
}

type Address struct {
	gorm.Model
	StaffID uint   `json:"userid"`
	Address string `json:"address"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
}
