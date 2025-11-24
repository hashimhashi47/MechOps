package models

import (
	"time"

	"gorm.io/gorm"
)

// user
type User struct {
	gorm.Model
	FirstName    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email" gorm:"unique"`
	Role         string    `json:"role"`
	Phone        string    `json:"phone"`
	Bookings     []Booking `gorm:"foreignKey:UserID"`
	Booked       []Booked  `gorm:"foreignKey:UserID"`
	Block        bool      `gorm:"default:false"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	RefreshToken string    `json:"-" gorm:"type:text"`
}

// RegisterBooking
type Booking struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey"`
	CarModel  string `json:"carmodel"`
	CarNumber string `json:"carnumber"`
	FuelType  string `json:"fueltype"`
	Problem   string `json:"problem"`
	Time      string `json:"time"`
	Date      string `json:"date"`
	Address   string `json:"Address"`
	LandMark  string `json:"landmark"`
	UserID    uint
}


//After user confirm service
type Booked struct {
	gorm.Model

	UserID           uint
	BookedID         string `json:"id" gorm:"primaryKey"`
	StaffPhoneNumber string `json:"staffphone"`
	Status           string `json:"status"`

	PaymentStatus string  `json:"paymentstatus"` //admin
	PaymentAmount float64 `json:"paymentamount"` //admin
	PaymentMode   string  `json:"paymentmode"`   //admin

	PickupTime *time.Time // setting admin
	PickedUpAt *time.Time // setting staff

	DeliveryTime *time.Time //  setting admin
	DeliveredAt  *time.Time // settig staff
}
