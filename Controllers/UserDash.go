package controllers

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	
	id, _ := c.Get("id")
	var user models.User

	userID := id.(uint)

	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Dashboard data",
		"user":    user,
	})

}
