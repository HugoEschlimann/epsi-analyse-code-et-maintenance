package controllers

import (
	"gin/models"
	"gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoanResources(c *gin.Context, db *gorm.DB, loans []*models.Loan) {
	if err := c.ShouldBindJSON(&loans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.LoanResources(c, db, loans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, "Loan(s) created successfully")
}