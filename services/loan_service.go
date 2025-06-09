package services

import (
	"errors"
	"gin/models"
	"strconv"
	"time"

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

		now := time.Now()
		loan.LoanDate = now.Format("02-01-2006")

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

func UpdateLoan(db *gorm.DB, id string) error {
	var loan models.Loan
	if err := db.First(&loan, id).Error; err != nil {
		return err
	}

	var resource models.Resource
	if err := db.First(&resource, loan.ResourceID).Error; err != nil {
		return err
	}

	resource.IsAvailable = true
	if err := db.Save(&resource).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLoan(db *gorm.DB, id string) error {
	var loan models.Loan
	loanId, _ := strconv.ParseUint(id, 10, 32)
	result := db.Where("id=?", loanId).Delete(&loan)
	return result.Error
}