package controllers

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	utils "MECHOPS/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Route to user page
func ManageUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "UsersPage.html", nil)
}

// Get all users data and show on Admin Pannel
func GetAllUsers(c *gin.Context) {
	var Users []models.User

	if err := db.DB.Find(&Users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to find users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Sucess": "Users find Successfully", "Users": Users})
}

// Upadate Users data
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var ExistingUser models.User

	if err := db.DB.Where("id = ?", id).First(&ExistingUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User not found"})
		return
	}

	var UpadatingUser models.User

	if err := c.ShouldBindJSON(&UpadatingUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if UpadatingUser.FirstName != "" {
		ExistingUser.FirstName = UpadatingUser.FirstName
	}

	if UpadatingUser.Email != "" {
		ExistingUser.Email = UpadatingUser.Email
	}

	if UpadatingUser.Phone != "" {
		ExistingUser.Phone = UpadatingUser.Phone
	}

	if UpadatingUser.Password != "" {
		hash, _ := utils.Hashing(UpadatingUser.Password)
		ExistingUser.Password = string(hash)
	}

	if err := db.DB.Save(&ExistingUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to update the user"})
		return
	}

	c.JSON(200, gin.H{"Success": "User Details updated successfully", "User": ExistingUser})
}

// Soft delete the user from database
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).Unscoped().Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to delete that user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "User deleted Succesfully"})
}

// Block / UnBlock user
func Blockuser(c *gin.Context) {
	id := c.Param("id")

	var Body struct {
		Block *bool `json:"block"`
	}

	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var User models.User

	if err := db.DB.Where("id = ?", id).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to find the user"})
		return
	}

	User.Block = *Body.Block

	if err := db.DB.Save(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to update block status"})
		return
	}

	status := "unblocked"
	if User.Block {
		status = "blocked"
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Block status updated " + status,
		"User":    User,
	})
}
