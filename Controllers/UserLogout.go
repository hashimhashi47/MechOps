package controllers

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// logout
func UserLogout(c *gin.Context) {
	id, _ := c.Get("id")
	userID := id.(uint)

	c.SetCookie("Token", "", -1, "/", "localhost", true, true)

	if err := db.DB.Model(models.User{}).Where("id = ?", userID).Update("refresh_token", "").Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to create refersh token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Sucees":"Logout suceesfully"})
}
