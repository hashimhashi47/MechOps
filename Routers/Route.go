package routers

import (
	constants "MECHOPS/Constants"
	controllers "MECHOPS/Controllers"
	middleware "MECHOPS/Middleware"
	services "MECHOPS/Services"
	"html/template"

	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {

	//Pusblic Routes
	Api := e.Group("/Api")

	{
		Api.POST("/SignUp", controllers.UserSignUp)
		Api.POST("/Login", controllers.UserLogin)
		Api.GET("/Admin/Logout", controllers.AdminLogout)
		Api.GET("/Logout", middleware.Middleware(constants.Admin, constants.Staff, constants.User), controllers.UserLogout)
	}


	//user Routes
	User := e.Group("/User")

	User.Use(middleware.Middleware(constants.User))
	{
		User.GET("/DashBoard", controllers.Dashboard)
		User.POST("/UpadteProfile", controllers.ProfileUpdate)
		User.POST("/BookService", services.Booking)
	}



	//Staff
	Staff := e.Group("/Staff")
	Staff.Use(middleware.Middleware(constants.Staff))
	{

	}

	//Admin
	Admin := e.Group("/Admin")

	{
		Admin.POST("/login", controllers.AdminLogin)
		Admin.GET("/login", controllers.AdminLogin)

		Admin.GET("/Dashboard", middleware.AdminAuth(), controllers.AdminDashboardPage)

		Admin.GET("/users/count", controllers.GetUsersCount)
		Admin.GET("/staff/count", controllers.GetStaffCount)
		Admin.GET("/bookings/count", controllers.GetBookingCount)
		Admin.GET("/bookings/next", controllers.GetNextService)
		Admin.GET("/bookings/recent", controllers.GetRecentBookings)

		Admin.GET("/Users", controllers.ManageUsers)

		Admin.GET("/Staff", controllers.ManageStaff)

		Admin.GET("/Bookings", controllers.ManageBookings)
		Admin.GET("/AllBooking",controllers.GetBookings)

		Admin.GET("/Slots", controllers.ManageSlots)
	}

	 // Load all HTML templates
    tmpl := template.Must(template.ParseGlob("Templates/*.html"))
    e.SetHTMLTemplate(tmpl)

}
