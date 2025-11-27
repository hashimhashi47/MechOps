package controllers

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "AdminDashboard.html", nil)
}


func GetUsersCount(c *gin.Context) {
    var count int64
    db.DB.Model(&models.User{}).Count(&count)
    c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetStaffCount(c *gin.Context) {
    var count int64
    db.DB.Model(&models.Staff{}).Count(&count)
    c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetBookingCount(c *gin.Context) {
    var count int64
    db.DB.Model(&models.Booking{}).Count(&count)
    c.JSON(http.StatusOK, gin.H{"count": count})
}


func GetNextService(c *gin.Context) {
    var booking models.Booking
    db.DB.Order("time asc").First(&booking)
    c.JSON(http.StatusOK, booking)
}

func GetRecentBookings(c *gin.Context) {
    var bookings []models.Booking
    db.DB.Order("created_at desc").Limit(5).Find(&bookings)
    c.JSON(http.StatusOK, bookings)
}