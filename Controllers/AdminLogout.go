package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//admin dashboard
func AdminLogout(c *gin.Context) {

	c.SetCookie("admin_id", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/Admin/login")
	
}
