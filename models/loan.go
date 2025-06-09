package models

type Loan struct {
	ID 		   uint   `gorm:"primaryKey" json:"id"`

	UserID     uint      `json:"user_id"`
	User       User      `gorm:"foreignKey:UserID"`

	ResourceID uint      `json:"resource_id"`
	Resource   Resource  `gorm:"foreignKey:ResourceID"`

	LoanDate   string 	 `json:"loan_date"`
	ReturnDate string 	 `json:"return_date"`
}