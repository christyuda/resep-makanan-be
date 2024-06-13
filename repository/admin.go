package repository

import (
	"Ayala-Crea/ResepBe/model"

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