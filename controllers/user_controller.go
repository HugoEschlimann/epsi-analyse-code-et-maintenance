package controllers

import (
	"fmt"
	"gin/logger"
	"gin/models"
	"gin/services"
	"gin/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nyaruka/phonenumbers"

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
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	logger.GetLogger().Info(fmt.Sprintf("Retrieved %d users successfully", len(users)))
	c.JSON(http.StatusOK, users)
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
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if err := utils.IsValidEmail(user.Email); !err {
		logger.GetLogger().Error(fmt.Sprintf("%s: %s", utils.ErrorEmail, user.Email))
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorEmail})
		return
	}

	phoneNumber, err := phonenumbers.Parse(user.Phone, "")
	if err != nil || !phonenumbers.IsValidNumber(phoneNumber) {
		logger.GetLogger().Error(fmt.Sprintf("%s: %s", utils.ErrorPhone, user.Phone))
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorPhone})
		return
	}

	user.PublicID = uuid.New()
	err = services.CreateUser(db, user)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	logger.GetLogger().Info(fmt.Sprintf("User with ID %d created successfully", user.ID))
	c.JSON(http.StatusCreated, user.PublicID.String())
}

// @Summary Update an existing user
// @Description Update an existing user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{uuid} [put]
func UpdateUser(c *gin.Context, db *gorm.DB, user *models.User) {
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if user.Email != "" && !utils.IsValidEmail(user.Email) {
		logger.GetLogger().Error(fmt.Sprintf("%s: %s", user.Email))
		c.JSON(http.StatusBadRequest, gin.H{"error": "%s"})
		return
	}

	if user.Phone != "" {
		phoneNumber, err := phonenumbers.Parse(user.Phone, "")
		if err != nil {
			logger.GetLogger().Error(fmt.Sprintf("%s: %s", utils.ErrorPhone, user.Phone))
			c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorPhone})
			return
		}
		if !phonenumbers.IsValidNumber(phoneNumber) {
			logger.GetLogger().Error(fmt.Sprintf("%s: %s", utils.ErrorPhone, user.Phone))
			c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorPhone})
			return
		}
	}
	
	uuidParam := c.Param("uuid")
	userUuid, err := uuid.Parse(uuidParam)
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("%s: %s", utils.ErrorUUID, uuidParam))
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorUUID})
		return
	}

	err = services.UpdateUser(db, userUuid.String(), user)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	logger.GetLogger().Info(fmt.Sprintf("User with UUID %s updated successfully", userUuid.String()))
	c.JSON(http.StatusOK, "User updated successfully")
}

// @Summary Restore a deleted user
// @Description Restore a deleted user by their UUID
// @Tags Users
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {string} string "User restored successfully"
// @Failure 500 {object} map[string]interface{}
// @Router /users/{uuid}/restore [patch]
func RestoreUser(c *gin.Context, db *gorm.DB) {
	uuidParam := c.Param("uuid")
	userUuid, err := uuid.Parse(uuidParam)
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("%s: %s", uuidParam))
		c.JSON(http.StatusBadRequest, gin.H{"error": "%s"})
		return
	}

	err = services.RestoreUser(db, userUuid.String())
	if err != nil {
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore user"})
		return
	}

	logger.GetLogger().Info(fmt.Sprintf("User with UUID %s restored successfully", userUuid.String()))
	c.JSON(http.StatusOK, gin.H{"message": "User restored successfully"})
}

// @Summary Archive a user
// @Description Archive a user by their UUID
// @Tags Users
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {string} string "User archive successfully"
// @Failure 500 {object} map[string]interface{}
// @Router /users/{uuid} [delete]
func ArchiveUser(c *gin.Context, db *gorm.DB) {
	uuidParam := c.Param("uuid")
	userUuid, err := uuid.Parse(uuidParam)
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("%s: %s", uuidParam))
		c.JSON(http.StatusBadRequest, gin.H{"error": "%s"})
		return
	}

	var nbLoans int64
	err = db.
		Model(&models.Loan{}).
		Where("user_uuid = ? AND return_date < ?", userUuid, time.Now().String()).
		Count(&nbLoans).Error

	if err != nil {
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count loans"})
		return
	}

	if nbLoans > 0 {
		logger.GetLogger().Error(fmt.Sprintf("User with UUID %s has active loans", userUuid.String()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "User has active loans and cannot be archived"})
		return
	}

	err = services.ArchiveUser(db, userUuid.String())
	if err != nil {
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to archive user"})
		return
	}

	logger.GetLogger().Info(fmt.Sprintf("User with UUID %s deleted successfully", userUuid.String()))
	c.JSON(http.StatusOK, gin.H{"message": "User archive successfully"})
}
