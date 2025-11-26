package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StaffLogin(c *gin.Context) {
	var Input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
}
