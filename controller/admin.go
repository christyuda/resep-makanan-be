package controller

import (
	repo "Ayala-Crea/ResepBe/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllReceipt(c *fiber.Ctx) error {
	// Cek token header autentikasi
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	db := c.Locals("db").(*gorm.DB)

	adminReceipt, err := repo.GetAllReceipeAdmin(db)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Terjadi Kesalahan dalam Get Data"})
	}

	// Jika tidak ada task yang ditemukan, mengembalikan pesan kesalahan
	if len(adminReceipt) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"code": http.StatusNotFound, "success": false, "status": "error", "message": "Data task tidak ditemukan", "data": nil})
	}

	// Jika tidak ada kesalahan, mengembalikan data task sebagai respons JSON
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    adminReceipt,
	}

	return c.Status(http.StatusOK).JSON(response)
}

// func CreateReciptAdmin(c *fiber.Ctx) error {

// }
