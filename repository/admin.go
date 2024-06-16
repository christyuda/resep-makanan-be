package repository

import (
	"Ayala-Crea/ResepBe/model"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

func GetAllReceipeAdmin(db *gorm.DB) ([]model.AdminReceipt, error) {
	var adminReceipt []model.AdminReceipt
	if err := db.
		Table("receipt").
		Select("receipt.recipe_id, users.id_user, users.username, receipt.title, receipt.description, receipt.ingredients, receipt.img").
		Joins("JOIN users ON receipt.id_user = users.id_user").
		Find(&adminReceipt).
		Error; err != nil {
		return nil, err
	}
	return adminReceipt, nil
}

func GetAllUsers(db *gorm.DB) ([]model.Users, error) {
	var users []model.Users
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func ValidateToken(tokenString string) (*model.JWTClaims, error) {
	secretKey := "secret_key" // Sesuaikan dengan kunci rahasia yang Anda gunakan untuk menandatangani token

	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}