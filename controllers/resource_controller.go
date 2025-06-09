package controllers

import (
	"fmt"
	"gin/models"
	"gin/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Create a new resource
// @Description Create a new resource with the provided details
// @Tags Resources
// @Accept json
// @Produce json
// @Param resource body models.Resource true "Resource details"
// @Success 201 {string} string "Resource created successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /resources [post]
func CreateResource(c *gin.Context, db *gorm.DB, resource *models.Resource) {
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateResource(db, resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, fmt.Sprintf("Resource %s created successfully", resource.Type))
}

// @Summary Update an existing resource
// @Description Update an existing resource with the provided details
// @Tags Resources
// @Accept json
// @Produce json
// @Param id path string true "Resource ID"
// @Param resource body models.Resource true "Resource details"
// @Success 200 {string} string "Resource updated successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /resources/{id} [put]
func UpdateResource(c *gin.Context, db *gorm.DB, resource *models.Resource) {
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	resourceId := c.Param("id")
	err := services.UpdateResource(db, resourceId, resource)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Resource %s not found", resourceId)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
    c.JSON(http.StatusOK, gin.H{"message": "Resource updated successfully"})

}

// @Summary Get all resources
// @Description Retrieve all resources
// @Tags Resources
// @Accept json
// @Produce json
// @Success 200 {array} models.Resource
// @Failure 500 {object} map[string]interface{}
// @Router /resources [get]
func GetResources(c *gin.Context, db *gorm.DB) {
	resources, err := services.GetResources(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resources)
}

// @Summary Delete a resource
// @Description Delete a resource by its ID
// @Tags Resources
// @Accept json
// @Produce json
// @Param id path string true "Resource ID"
// @Success 200 {string} string "Resource deleted successfully"
// @Failure 500 {object} map[string]interface{}
// @Router /resources/{id} [delete]
func DeleteResource(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	err := services.DeleteResource(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}