package controllers

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Route to admindashboard
func AdminDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "AdminDashboard.html", nil)
}

//get count of users
func GetUsersCount(c *gin.Context) {
	var count int64
	db.DB.Model(&models.User{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

//get count of staff
func GetStaffCount(c *gin.Context) {
	var count int64
	db.DB.Model(&models.Staff{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"count": count})
}


//get count of booking
func GetBookingCount(c *gin.Context) {
	var count int64
	db.DB.Model(&models.Booking{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

//get booking details
func GetNextService(c *gin.Context) {
	var booking models.Booking
	db.DB.First(&booking)
	c.JSON(http.StatusOK, booking)
}


//get % booking details to show dashboard
func GetRecentBookings(c *gin.Context) {
	var bookings []models.Booking
	db.DB.Limit(5).Find(&bookings)
	c.JSON(http.StatusOK, bookings)
}
