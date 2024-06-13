package main

import (
	"Ayala-Crea/ResepBe/config"
	"Ayala-Crea/ResepBe/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Membuat aplikasi Fiber
	app := fiber.New()

	// Koneksi ke database
	db := config.CreateDBConnection()

	// Middleware untuk logging
	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	// Middleware untuk mengatur CORS
	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Menyimpan koneksi database dalam context Fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Menetapkan rute untuk handler buku
	routes.SetupTaskRoutes(app)

	// Menambahkan rute untuk menyajikan file statis dari direktori "images"
	app.Static("/img", "/img")

	// Menambahkan rute untuk menguji apakah server berjalan
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running. Go to /img/<filename> to view the image.")
	})

	// Menjalankan server Fiber pada port 3000
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
