package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := c.Cookie("admin_id")
		if err != nil || adminID == "" {
			c.Redirect(http.StatusSeeOther, "/Admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}