package services

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	utils "MECHOPS/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserBooking
func Booking(c *gin.Context) {

	UserID := c.MustGet("id").(uint)

	var Input struct {
		CarModel  string `json:"carmodel" binding:"required"`
		CarNumber string `json:"carnumber" binding:"required"`
		FuelType  string `json:"fueltype" binding:"required"`
		Problem   string `json:"problem" binding:"required"`
		Time      string `json:"time" binding:"required"`
		Date      string `json:"date" binding:"required"`
		Address   string `json:"address" binding:"required"`
		LandMark  string `json:"landmark" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ID := utils.RandomIDGenerate("BOOK")

	Booking := models.Booking{
		UserID:    UserID,
		ID:        ID,
		CarModel:  Input.CarModel,
		CarNumber: Input.CarNumber,
		FuelType:  Input.FuelType,
		Problem:   Input.Problem,
		Time:      Input.Time,
		Date:      Input.Date,
		Address:   Input.Address,
		LandMark:  Input.LandMark,
	}

	if err := db.DB.Create(&Booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to Booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Sucess": "Our Team will contact soon",
		"ID":     Booking.ID,
	})
}
