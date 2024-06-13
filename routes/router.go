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

	// app.Get("/tasks", controller.GetAllTask)
	// app.Get("/task/get", controller.GetTaskById)
	// app.Get("/task/taskuser", controller.GetTaskByIdUser)
	// app.Post("/task/insert", controller.InsertTask)
	// app.Put("/task/update", controller.UpdateTask)
	// app.Delete("/task/delete", controller.DeleteTask)

	// app.Get("/roles", controller.GetAllRole)
	// app.Get("/role/get/:id_role", controller.GetRoleById)
	// app.Post("/role/insert", controller.InsertRole)
	// app.Put("/role/update/:id_role", controller.UpdateRole)
	// app.Delete("/role/delete/:id_role", controller.DeleteRole)
}
