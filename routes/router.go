package routes

import (
	"Ayala-Crea/ResepBe/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App) {
	app.Post("/register", controller.RegisterUser)
	app.Post("/login", controller.LoginUser)

	// Endpoint untuk cek sudah login atau belum
	app.Get("/getme", controller.GetMe)

	app.Post("/receipe", controller.CreateReceipe)
	app.Get("/receipe", controller.GetAllReceipe)
	app.Get("receipe/get", controller.GetReceipeById)
	app.Put("/receipe/update", controller.UpdateReceipe)
	app.Delete("receipe/delete", controller.DeleteReceipt)
	app.Get("/resep/user", controller.GetReceiptByUser)

	app.Get("/users", controller.GetAllUser)
	app.Post("/users", controller.CreateUser)
	app.Get("/users/Get", controller.GetByIdUser)

}
