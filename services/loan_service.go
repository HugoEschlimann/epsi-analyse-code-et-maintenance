package services

import (
	"errors"
	"gin/models"

	"gorm.io/gorm"
)

func LoanResources(db *gorm.DB, loans []*models.Loan) error {
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

		loan.LoanDate = "CURRENT_TIMESTAMP" // Assuming you want to set the current timestamp

		if err := db.Create(loan).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetLoans(db *gorm.DB) ([]models.Loan, error) {
	var loans []models.Loan
	if err := db.Preload("User").Preload("Resource").Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}

// update is_available resource field and delete loan
func RestituteResources(db *gorm.DB, loans []*models.Loan) error {
	for _, loan := range loans {
		var resource models.Resource
		if err := db.First(&resource, loan.ResourceID).Error; err != nil {
			return err
		}

		resource.IsAvailable = true
		if err := db.Save(&resource).Error; err != nil {
			return err
		}

		if err := db.Where("user_id = ? AND resource_id = ?", loan.UserID, loan.ResourceID).Delete(&models.Loan{}).Error; err != nil {
			return err
		}
	}
	return nil
}