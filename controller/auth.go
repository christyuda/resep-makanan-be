package controller

import (
	"Ayala-Crea/ResepBe/model"
	repo "Ayala-Crea/ResepBe/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	var user model.Users

	db := c.Locals("db").(*gorm.DB)

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request Body Invalid",
		})
	}

	// Menyimpan user ke database menggunakan repository
	if err := repo.CreateUser(db, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Register",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil Register!",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var user model.Users
	db := c.Locals("db").(*gorm.DB)

	// Parsing request body into user struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request data",
		})
	}

	// Validate email or username provided
	if user.Email == "" && user.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request must include email or username",
		})
	}

	// Retrieve user data from the database
	var userData *model.Users
	var err error

	if user.Email != "" {
		userData, err = repo.GetUserByEmail(db, user.Email)
	} else if user.Username != "" {
		userData, err = repo.GetUserByUsername(db, user.Username)
	}

	if err != nil || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username or password incorrect",
		})
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username or password incorrect",
		})
	}

	// Generate JWT token
	token, err := repo.GenerateToken(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	// Get role information
	role, err := repo.GetUserByRoleId(db, userData.IdRole)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve role",
		})
	}

	// Return successful response with token and role data
	return c.JSON(fiber.Map{
		"token": token,
		"role":  role,
	})
}

func GetMe(c *fiber.Ctx) error {
	// Mendapatkan token JWT dari header
	tokenString := c.Get("login")
	if tokenString == "" {
		return fiber.NewError(fiber.StatusNotFound, "Token tidak ditemukan di header")
	}

	// Parse token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return err
	}

	// Memeriksa apakah token JWT valid
	if !token.Valid {
		return fiber.NewError(fiber.StatusBadRequest, "Token Invalid")
	}

	// Mengekstrak klaim dari token JWT
	claims := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id_user"].(float64)) // Mengambil nilai id_user dan mengonversinya ke uint
	db := c.Locals("db").(*gorm.DB)

	// Cari user di database berdasarkan user ID menggunakan repository
	userData, err := repo.GetUserById(db, userId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"user": userData,
	})
}
