package services

import (
	"errors"
	"gin/models"
	"strconv"

	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	db.Omit("Password").Find(&users)
	return users, nil
}

func GetUserById(db *gorm.DB, id string) (*models.User, error) {
	var user models.User
	db.Omit("Password").Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func CreateUser(db *gorm.DB, user *models.User) error {
	if err := user.HashPassword(); err != nil {
		return err
	}

	result := db.Create(user)
	return result.Error
}

func UpdateUser(db *gorm.DB, user *models.User) error {
	db.Model(&user).Updates(user)
	return nil
}

func DeleteUser(db *gorm.DB, id string) error {
	var user models.User
	userId, _ := strconv.ParseUint(id, 10, 32)
	db.Where("id=?", userId).Delete(&user)
	return nil
}