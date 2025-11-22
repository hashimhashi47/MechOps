package routers

import (
	controllers "MECHOPS/Controllers"

	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	
	//user Routes
	User := e.Group("/User")
	{
		User.POST("/SignUp", controllers.UserSignUp)
		User.POST("/Login", controllers.UserLogin)
	}
}
