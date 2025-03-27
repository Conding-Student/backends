package routes

import (
	//"intern_template_v1/controller"
	controller "intern_template_v1/controller/Admin"
	Usercontroller "intern_template_v1/controller/auth"
	landlordcontroller "intern_template_v1/controller/landlord"
	"intern_template_v1/middleware"

	"github.com/gofiber/fiber/v2"
	//"golang.org/x/crypto/nacl/auth"
)

func AppRoutes(app *fiber.App) {
	// SAMPLE ENDPOINT
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Golang World!")
	})

	// CREATE YOUR ENDPOINTS HERE
	//app.Get("/try", controller.SampleController2)
	//app.Get("/try1", controller.SampleController1)

	//app.Post("/create", controller.UserRegistration)
	//app.Post("/read", controller.ReadUser)
	//app.Post("/create", controller.CreateBook)
	//app.Get("/get/all/books", controller.GetAllBooks)
	//app.Get("/get/books/:id", controller.Getbook)
	//app.Put("/update/book/:id", controller.UpdateBook)
	//app.Post("/register/user", controller.RegisterUser)
	//app.Get("/get/user/:id", controller.GetUser)
	//app.Get("/get/all/user", controller.GetAllUsers)
	app.Post("/registertenant/account", Usercontroller.RegisterTenant)
	app.Post("/registerlandlord/account", Usercontroller.RegisterLandlord)
	app.Post("/loginuser/account", Usercontroller.LoginUser)
	//app.Post("/addrentallisting", landlordcontroller.CreateApartment)
	app.Post("/property/add", middleware.AuthMiddleware, landlordcontroller.CreateApartment)
	app.Get("/property/get", middleware.AuthMiddleware, landlordcontroller.FetchApartmentsByLandlord)
	app.Post("/admin/register", controller.RegisterAdmin)
	app.Post("/admin/login", controller.LoginHandler)

}
