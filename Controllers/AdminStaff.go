package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ManageStaff(c *gin.Context) {
	c.HTML(http.StatusOK, "StaffPage.html", nil)
}
