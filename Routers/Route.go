package routers

import (
	constants "MECHOPS/Constants"
	controllers "MECHOPS/Controllers"
	middleware "MECHOPS/Middleware"

	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {

	//user Routes
	User := e.Group("/User")
	{
		User.POST("/SignUp", controllers.UserSignUp)
		User.POST("/Login", controllers.UserLogin)
		User.GET("/DashBoard",middleware.Middleware(constants.User),controllers.Dashboard)
	}
	
}
