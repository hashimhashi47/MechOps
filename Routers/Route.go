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

		//Dashboard routes and endpoints
		Admin.GET("/Dashboard", middleware.AdminAuth(), controllers.AdminDashboardPage)
		Admin.GET("/users/count", controllers.GetUsersCount)
		Admin.GET("/staff/count", controllers.GetStaffCount)
		Admin.GET("/bookings/count", controllers.GetBookingCount)
		Admin.GET("/bookings/next", controllers.GetNextService)
		Admin.GET("/bookings/recent", controllers.GetRecentBookings)

		//User routes and endpoints
		Admin.GET("/Users", controllers.ManageUsers)
		Admin.GET("/GetAllusers", controllers.GetAllUsers)
		Admin.PUT("/UpdateUsers/:id", controllers.UpdateUser)
		Admin.POST("/AddUser", controllers.UserSignUp)
		Admin.DELETE("/Delete/:id", controllers.DeleteUser)
		Admin.PUT("User/Block/:id", controllers.Blockuser)

		//Staff routes and endPoints
		Admin.GET("/Staff", controllers.ManageStaff)

		//bookings Management and endpoints
		Admin.GET("/Bookings", controllers.ManageBookings)
		Admin.GET("/AllBooking", controllers.GetBookings)
		Admin.PUT("/AllBooking/update/:id", controllers.UpdateBooking)

		//slot management and end points
		Admin.GET("/Slots", controllers.ManageSlots)

	}

	// Load all HTML templates
	tmpl := template.Must(template.ParseGlob("Templates/*.html"))
	e.SetHTMLTemplate(tmpl)

}
