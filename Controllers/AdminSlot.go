package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ManageSlots(c *gin.Context) {
	c.HTML(http.StatusOK, "SlotOnGarage.html", nil)
}
