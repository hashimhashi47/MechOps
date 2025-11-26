package controllers

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)


//admin Login
func AdminLogin(c *gin.Context) {

	Email := c.PostForm("email")
	Password := c.PostForm("password")

	var admin models.Admin

	result:= db.DB.Where("email = ?", Email).First(&admin)


	if result.Error != nil {
		c.HTML(http.StatusUnauthorized, "Login.html", gin.H{
			"error": "Admin not found",
		})
		return
	}

	if Password != admin.Password {
		c.HTML(http.StatusUnauthorized, "Login.html", gin.H{
			"error": "Invalid credentials",
		})
		return
	}


	c.SetCookie("admin_id", admin.Name, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/Admin/Dashboard")
}
