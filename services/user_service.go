package services

import (
	"gin/models"
	"strconv"

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


func UpdateUser(db *gorm.DB, id string, user *models.User) error {
	if err := db.First(&models.User{}, id).Error; err != nil {
		return err
	}

	userId, _ := strconv.ParseUint(id, 10, 32)
	result := db.Model(&models.User{}).Where("id = ?", userId).Updates(user)
	return result.Error
}

func DeleteUser(db *gorm.DB, id string) error {
	var user models.User
	userId, _ := strconv.ParseUint(id, 10, 32)
	result := db.Where("id=?", userId).Delete(&user)
	return result.Error
}