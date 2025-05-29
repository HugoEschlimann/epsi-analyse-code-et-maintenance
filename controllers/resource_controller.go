package controllers

import (
	"fmt"
	"gin/models"
	"gin/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateResource(c *gin.Context, db *gorm.DB, resource *models.Resource) {
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateResource(db, resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, fmt.Sprintf("Resource %s created successfully", resource.Type))
}

func UpdateResource(c *gin.Context, db *gorm.DB, resource *models.Resource) {
	resourceId := c.Param("id")

	// if err := c.ShouldBindJSON(&resource); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	err := services.UpdateResource(c, db, resourceId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
    c.JSON(http.StatusOK, gin.H{"message": "Resource updated successfully", "resource": resource})

}

func GetResources(c *gin.Context, db *gorm.DB) {
	resources, err := services.GetResources(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resources)
}

func DeleteResource(c *gin.Context, db *gorm.DB, id string) {
	err := services.DeleteResource(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}