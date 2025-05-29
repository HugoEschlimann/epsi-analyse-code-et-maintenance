package services

import (
	"gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateResource(db *gorm.DB, resource *models.Resource) error {
	result := db.Create(resource)
	return result.Error
}

func GetResources(db *gorm.DB) ([]models.Resource, error) {
	var resources []models.Resource
	db.Omit("Password").Find(&resources)
	return resources, nil
}

func UpdateResource(c *gin.Context, db *gorm.DB, resourceId string) error {
	var resource models.Resource

	if err := db.First(&resource, resourceId).Error; err != nil {
		return err
	}

	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	if err := db.Model(&resource).Updates(resource).Error; err != nil {
		return err
	}
	
	return nil
}

func DeleteResource(db *gorm.DB, id string) error {
	var resource models.Resource
	resourceId, _ := strconv.ParseUint(id, 10, 32)
	db.Where("id=?", resourceId).Delete(&resource)
	return nil
}