package controllers

import (
	"gin/models"
	"gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Loan resources
// @Description Loan resources to users
// @Tags Loans
// @Accept json
// @Produce json
// @Param loans body []models.Loan true "List of loans"
// @Success 201 {string} string "Loan(s) created successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /loans [post]
func LoanResources(c *gin.Context, db *gorm.DB, loans []*models.Loan) {
	if err := c.ShouldBindJSON(&loans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.LoanResources(db, loans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Loan(s) created successfully")
}

// @Summary Get all loans
// @Description Retrieve all loans
// @Tags Loans
// @Produce json
// @Success 200 {array} models.Loan
// @Failure 500 {object} map[string]interface{}
// @Router /loans [get]
func GetLoans(c *gin.Context, db *gorm.DB) {
	loans, err := services.GetLoans(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loans)
}

// @Summary Restitute resources
// @Description Restitute resources from loans
// @Tags Loans
// @Accept json
// @Produce json
// @Param loans body []models.Loan true "List of loans to restitute"
// @Success 200 {string} string "Resources restituted successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /restitute [post]
func Restitute(c *gin.Context, db *gorm.DB, loans []*models.Loan) {
	if err := c.ShouldBindJSON(&loans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.RestituteResources(db, loans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Resource(s) restituted successfully")
}