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
    result := db.DB.Order("created_at desc").Find(&bookings)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    if bookings == nil {
        bookings = []models.Booking{}
    }

    c.JSON(http.StatusOK, bookings)
}
