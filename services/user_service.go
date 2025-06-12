package services

import (
	"gin/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	result := db.Create(user)
	return result.Error
}

func GetUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	db.Find(&users)
	return users, nil
}


func UpdateUser(db *gorm.DB, userUuid string, user *models.User) error {
	var currentUser models.User
	if err := db.First(&currentUser, "public_id = ?", userUuid).Error; err != nil {
		return err
	}

	result := db.Model(&models.User{}).Where("id = ?", currentUser.ID).Updates(user)
	return result.Error
}

func DeleteUser(db *gorm.DB, uuid string) error {
	var user models.User
	result := db.Where("public_id=?", uuid).Delete(&user)
	return result.Error
}