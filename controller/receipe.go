package controller

import (
	"Ayala-Crea/ResepBe/model"
	repo "Ayala-Crea/ResepBe/repository"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReceipe(c *fiber.Ctx) error {
	// cek token header auth
	tokenStr := c.Get("login")
	if tokenStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// parse untuk mendapatkan id
	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil // Ganti "secret_key" dengan kunci rahasia Anda
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	idUser := claims.IdUser

	var receipe model.Receipt

	// Ambil data form
	if err := c.BodyParser(&receipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Ambil file image
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Gagal mengambil file"})
	}

	// Simpan file image
	imagePath := fmt.Sprintf("./img/%d_%s", time.Now().UnixNano(), file.Filename)
	if err := c.SaveFile(file, imagePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan file"})
	}

	// Set path image di model
	receipe.Image = imagePath
	receipe.IdUser = int(idUser)

	db := c.Locals("db").(*gorm.DB)

	if err := repo.InsertReceipt(db, receipe); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"code": http.StatusCreated, "success": true, "status": "success", "message": "Task berhasil disimpan", "data": receipe})
}

func GetAllReceipe(c *fiber.Ctx) error {
	// cek token header auth
	tokenStr := c.Get("login")
	if tokenStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// parse untuk mendapatkan id
	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil // Ganti "secret_key" dengan kunci rahasia Anda
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	idUser := claims.IdUser
	var receipes model.Receipt
	db := c.Locals("db").(*gorm.DB)

	receipe, err := repo.GetAllReceipe(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data role, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	receipes.IdUser = int(idUser)

	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    receipe,
	}

	return c.Status(http.StatusOK).JSON(response)
}
func GetReceipeById(c *fiber.Ctx) error {
	// Cek token header autentikasi
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// mencari id dari data
	id := c.Query("recipe_id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID task tidak boleh kosong"})
	}

	db := c.Locals("db").(*gorm.DB)

	receipe, err := repo.GetReceipetById(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Data Tidak Ditemukan"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": receipe})
}

func UpdateReceipe(c *fiber.Ctx) error {
	// Cek token header autentikasi
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(http.StatusBadRequest, "Header tidak ada")
	}

	// Mengambil Parameter ID
	id := c.Query("recipe_id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID task tidak ditemukan"})
	}

	var updateRecipe model.Receipt

	if err := c.BodyParser(&updateRecipe); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	db := c.Locals("db").(*gorm.DB)

	existingRecipe, err := repo.GetReceipetById(db, id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ID task tidak ditemukan"})
	}

	// Ambil file image jika ada
	file, err := c.FormFile("image")
	if err == nil {
		// Simpan file image baru
		imagePath := fmt.Sprintf("./img/%d_%s", time.Now().UnixNano(), file.Filename)
		if err := c.SaveFile(file, imagePath); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan file"})
		}
		updateRecipe.Image = imagePath
	} else {
		// Jika tidak ada file baru, gunakan path gambar lama
		updateRecipe.Image = existingRecipe.Image
	}

	if err := repo.UpdateReceipt(db, id, updateRecipe); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui receipt"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Receipt berhasil diperbarui"})
}

func DeleteReceipt(c *fiber.Ctx) error {
	// Cek token header autentikasi
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(http.StatusBadRequest, "Header tidak ada")
	}

	id := c.Query("recipe_id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID task tidak ditemukan"})
	}

	db := c.Locals("db").(*gorm.DB)

	receipt, err := repo.GetReceipetById(db, id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ID task tidak ditemukan"})
	}

	if err := repo.DeleteReceipt(db, receipt.IdReceipe); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal Menghapus Data"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Receipt berhasil dihapus"})
}

func GetReceiptByUser(c *fiber.Ctx) error {
	// Cek token header autentikasi
	tokenStr := c.Get("login")
	if tokenStr == "" {
		return fiber.NewError(http.StatusBadRequest, "Header tidak ada")
	}

	// Mendapatkan id user dari token
	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil // Ganti "secret_key" dengan kunci rahasia Anda
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	idUser := claims.IdUser
	db := c.Locals("db").(*gorm.DB)

	receipts, err := repo.GetReceiptByUser(db, idUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": receipts})
}
