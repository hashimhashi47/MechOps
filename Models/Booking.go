package models

import (
	"time"

	"gorm.io/gorm"
)

// RegisterBooking
type Booking struct {
	ID        string `json:"id" gorm:"primaryKey"`
	CarModel  string `json:"car"`
	CarNumber string `json:"carnumber"`
	FuelType  string `json:"fueltype"`
	Problem   string `json:"service"`
	Time      string `json:"time"`
	Date      string `json:"date"`
	Address   string `json:"Address"`
	LandMark  string `json:"landmark" `
	Status    string `json:"status"`
	UserID    uint
}

// After user confirm service
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
