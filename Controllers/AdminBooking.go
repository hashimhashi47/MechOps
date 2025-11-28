package controllers

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ManageBookings(c *gin.Context) {
	c.HTML(http.StatusOK, "BookingManagement.html", nil)
}

func GetBookings(c *gin.Context) {

	var bookings []models.Booking

	result := db.DB.Find(&bookings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func UpdateBooking(c *gin.Context) {
	id := c.Param("id")

	var booking models.Booking
	if err := db.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to find booking"})
		return
	}

	var updateData struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if updateData.Status != "" {
		booking.Status = updateData.Status
	}

	if err := db.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Booking updated successfully", "data": booking})
}
