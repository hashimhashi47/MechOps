package controllers

import (
	constants "MECHOPS/Constants"
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	utils "MECHOPS/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var Input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	//Check the user is exisiting
	var User models.User
	if err := db.DB.Where("email = ?", Input.Email).First(&User).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "User not found"})
		return
	}
	fmt.Println(User)

	if User.Block && User.Role == constants.User {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "your account is blocked"})
		return
	}

	err := utils.HashCompare(User.Password, Input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid password"})
		return
	}

	//TOKENS
	AccessToken, err := utils.AccessToken(User.ID, User.Email, User.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to issue AccessToken"})
		return
	}

	RefershToken, err := utils.RefershToken(User.ID, User.Email, User.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to issue RefershToken"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Sucess":  "User Logged In Sucessfully",
		"Access":  AccessToken,
		"Refersh": RefershToken,
		"role":    User.Role,
	})
}
