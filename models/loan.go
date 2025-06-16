package models

import "github.com/google/uuid"

type Loan struct {
	ID 		   uint   		`gorm:"primaryKey;constraint:OnDelete:RESTRICT;" json:"id"`

	UserUUID   uuid.UUID    `json:"user_uuid"`
	User       User      	`gorm:"foreignKey:UserUUID"`

	ResourceID uint      	`json:"resource_id"`
	Resource   Resource  	`gorm:"foreignKey:ResourceID"`

	LoanDate   string 	 	`json:"loan_date"`
	ReturnDate string 	 	`json:"return_date"`
}