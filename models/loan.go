package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model

	UserID     uint      `json:"userid"`
	User       User      `gorm:"foreignKey:UserID"`

	ResourceID uint      `json:"resourceid"`
	Resource   Resource  `gorm:"foreignKey:ResourceID"`

	LoanDate   *time.Time `gorm:"loandate"`
	ReturnDate *time.Time `json:"returndate"`
}