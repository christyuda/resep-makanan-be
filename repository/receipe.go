package repository

import (
	"Ayala-Crea/ResepBe/model"

	"gorm.io/gorm"
)

func InsertReceipt(db *gorm.DB, receipe model.Receipt) error {
	if err := db.Create(&receipe).Error; err != nil {
		return err
	}
	return nil
}

func GetAllReceipe(db *gorm.DB) ([]model.Receipt, error) {
	var receipe []model.Receipt
	if err := db.Find(&receipe).Error; err != nil {
		return nil, err
	}
	return receipe, nil
}

func GetReceipetById(db *gorm.DB, id string) (model.Receipt, error) {
	var receipe model.Receipt
	if err := db.First(&receipe, "recipe_id = ?", id).Error; err != nil {
		return receipe, err
	}
	return receipe, nil
}

func UpdateReceipt(db *gorm.DB, id string, updateReceipt model.Receipt) error {
	if err := db.Model(&model.Receipt{}).Where("recipe_id = ?", id).Updates(updateReceipt).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReceipt(db *gorm.DB, id uint) error {
	if err := db.Delete(&model.Receipt{}, id).Error; err != nil {
		return err
	}
	return nil
}

func GetReceiptByUser(db *gorm.DB, id uint) ([]model.Receipt, error) {
	var receipts []model.Receipt
	if err := db.Find(&receipts, "id_user = ?", id).Error; err != nil {
		return nil, err
	}
	return receipts, nil
}
