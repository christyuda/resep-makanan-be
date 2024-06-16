package controller

import (
	"Ayala-Crea/ResepBe/model"
	repo "Ayala-Crea/ResepBe/repository"
	"net/http"
	"strconv"

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

func GetAllUser(c *fiber.Ctx) error {
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// Validasi token dan ambil informasi pengguna dari token
	claims, err := repo.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak valid",
		})
	}

	// Cek apakah pengguna memiliki id_role 1 (admin)
	if claims.IdRole != 1 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Anda tidak memiliki akses untuk melihat semua pengguna",
		})
	}

	db := c.Locals("db").(*gorm.DB)

	dataUsers, err := repo.GetAllUsers(db)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Terjadi Kesalahan dalam Get Data"})
	}
	// Jika tidak ada task yang ditemukan, mengembalikan pesan kesalahan
	if len(dataUsers) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"code": http.StatusNotFound, "success": false, "status": "error", "message": "Data task tidak ditemukan", "data": nil})
	}

	// Jika tidak ada kesalahan, mengembalikan data task sebagai respons JSON
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    dataUsers,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func CreateUser(c *fiber.Ctx) error {
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// Validasi token dan ambil informasi pengguna dari token
	claims, err := repo.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak valid",
		})
	}

	// Cek apakah pengguna memiliki id_role 1 (admin)
	if claims.IdRole != 1 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Anda tidak memiliki akses untuk melihat semua pengguna",
		})
	}

	var user model.Users

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request Body Invalid",
		})
	}

	db := c.Locals("db").(*gorm.DB)

	// Menyimpan user ke database menggunakan repository
	if err := repo.CreateUser(db, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Register",
		})
	}

	// Mengembalikan respons dengan data pengguna yang baru dibuat
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil Register!",
		"user": fiber.Map{
			"id_user":  user.IdUser,
			"nama":     user.Nama,
			"username": user.Username,
			"email":    user.Email,
			"id_role":  user.IdRole,
			// Tambahkan detail lain yang perlu ditampilkan
		},
	})
}

func GetByIdUser(c *fiber.Ctx) error {
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// Validasi token dan ambil informasi pengguna dari token
	claims, err := repo.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak valid",
		})
	}

	// Ambil `id_user` dari query parameter
	idParam := c.Params("id_user")
	if idParam == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID user tidak boleh kosong"})
	}

	// Konversi idParam ke uint
	idUser, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID user tidak valid",
		})
	}

	// Cek apakah pengguna memiliki id_role 1 (admin) atau meminta data dirinya sendiri
	if claims.IdRole != 1 && claims.IdUser != uint(idUser) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Anda tidak memiliki akses untuk melihat data pengguna ini",
		})
	}

	db := c.Locals("db").(*gorm.DB)

	// Ambil data pengguna dari database
	dataUser, err := repo.GetUserById(db, uint(idUser))
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Data tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": dataUser})
}
