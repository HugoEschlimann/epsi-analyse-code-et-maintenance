package services

import (
	"gin/models"
	"strconv"

	"gorm.io/gorm"
)

func CreateResource(db *gorm.DB, resource *models.Resource) error {
	result := db.Create(resource)
	return result.Error
}

func GetResources(db *gorm.DB) ([]models.Resource, error) {
	var resources []models.Resource
	db.Find(&resources)
	return resources, nil
}

func UpdateResource(db *gorm.DB, id string, resource *models.Resource) error {
	if err := db.First(&models.Resource{}, id).Error; err != nil {
		return err
	}

	resourceId, _ := strconv.ParseUint(id, 10, 32)
	result := db.Model(&models.Resource{}).Where("id = ?", resourceId).Updates(resource)
	return result.Error
}

func DeleteResource(db *gorm.DB, id string) error {
	var resource models.Resource
	resourceId, _ := strconv.ParseUint(id, 10, 32)
	result := db.Where("id=?", resourceId).Delete(&resource)
	return result.Error
}