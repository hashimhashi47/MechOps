package controllers

import (
	constants "MECHOPS/Constants"
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	utils "MECHOPS/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp(c *gin.Context) {

	//validation
	var Input struct {
		FirstName string `json:"firstname" binding:"required"`
		Lastname  string `json:"lastname" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Phone     string `json:"phone" binding:"required,min=10"`
		Password  string `json:"password" binding:"required,min=6,max=16"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}

	//existing user
	var Exist models.User
	if err := db.DB.Where("email = ?", Input.Email).First(&Exist).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Email already existing.."})
		return
	}

	if err := db.DB.Where("phone = ?", Input.Phone).First(&Exist).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Phone Number already existing.."})
		return
	}

	//hashing
	Hash, err := utils.Hashing(Input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to HashPassword"})
		return
	}

	CreateUser := models.User{
		FirstName: Input.FirstName,
		Lastname:  Input.Lastname,
		Email:     Input.Email,
		Password:  Hash,
		Phone:     Input.Phone,
		Role:      constants.User,
	}

	//Add user on db
	if err := db.DB.Create(&CreateUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to Create User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "SignUp Sucessfully"})
}
