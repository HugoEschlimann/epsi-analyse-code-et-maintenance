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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	err := services.LoanResources(db, loans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get loans"})
		return
	}
	c.JSON(http.StatusOK, loans)
}

// @Summary Update an existing loan
// @Description Update an existing loan with the provided details
// @Tags Loans
// @Accept json
// @Produce json
// @Param id path string true "Loan ID"
// @Param loan body models.Loan true "Loan details"
// @Success 200 {string} string "Loan updated successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /loans/{id} [put]
func UpdateLoan(c *gin.Context, db *gorm.DB) {
	loanId := c.Param("id")
	err := services.UpdateLoan(db, loanId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Loan updated successfully"})
}

// @Summary Restitute resources
// @Description Restitute resources from loans
// @Tags Loans
// @Accept json
// @Produce json
// @Param loans body models.Loan true "Loan to delete"
// @Success 200 {string} string "Resources restituted successfully"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /restitute [post]
func DeleteLoan(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	err := services.DeleteLoan(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete loan"})
		return
	}
	c.JSON(http.StatusOK, "Loan deleted successfully")
}
