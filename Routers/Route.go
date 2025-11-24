package routers

import (
	constants "MECHOPS/Constants"
	controllers "MECHOPS/Controllers"
	middleware "MECHOPS/Middleware"

	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {

	//Pusblic Routes
	Api := e.Group("/Api")
	{
		Api.POST("/SignUp", controllers.UserSignUp)
		Api.POST("/Login", controllers.UserLogin)
		Api.GET("/Logout", middleware.Middleware(constants.Admin, constants.Staff, constants.User), controllers.UserLogout)
	}

	//user Routes
	User := e.Group("/User")
	User.Use(middleware.Middleware(constants.User))
	{
		User.GET("/DashBoard", controllers.Dashboard)
		User.POST("/UpadteProfile", controllers.ProfileUpdate)
	}

}
