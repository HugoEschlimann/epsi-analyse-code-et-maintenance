package services

import (
	"errors"
	"gin/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoanResources(c *gin.Context, db *gorm.DB, loans []*models.Loan) error {
	for _, loan := range loans {
		var resource models.Resource
		if err := db.First(&resource, loan.ResourceID).Error; err != nil {
			return err
		}

		if !resource.IsAvailable {
			return errors.New("resource is not available")
		}

		resource.IsAvailable = false
		if err := db.Save(&resource).Error; err != nil {
			return err
		}

		now := time.Now()
		loan.LoanDate = &now

		if err := db.Create(loan).Error; err != nil {
			return err
		}
	}
	return nil
}