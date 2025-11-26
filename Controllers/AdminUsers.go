package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ManageUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "UsersPage.html", nil)
}
