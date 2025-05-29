package controllers

import (
	"fmt"
	"gin/models"
	"gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary List all users
// @Description Get a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func GetUsers(c *gin.Context, db *gorm.DB) {
	users, err := services.GetUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [get]
func GetUserById(c *gin.Context, db *gorm.DB, id string) {
	user, err := services.GetUserById(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *gin.Context, db *gorm.DB, user *models.User) {
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("User %s created successfully", user.Name))
}

// @Summary Update an existing user
// @Description Update an existing user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [put]
func UpdateUser(c *gin.Context, db *gorm.DB, user *models.User) {
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.UpdateUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("User %s updated successfully", user.Name))
}

// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context, db *gorm.DB, id string) {
	err := services.DeleteUser(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
