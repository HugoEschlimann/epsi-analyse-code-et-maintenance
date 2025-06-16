package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uint      		`gorm:"primaryKey" json:"id"`
	PublicID    uuid.UUID 		`json:"public_id"`
	Lastname    string    		`json:"lastname"`
	Firstname   string    		`json:"firstname"`
	Email       string    		`gorm:"unique" json:"email"`
	Phone       string    		`gorm:"unique" json:"phone"`
	Nationality string    		`json:"nationality"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
