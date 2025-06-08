package models

type User struct {
	ID 			 uint   `gorm:"primaryKey" json:"id"`
    Lastname     string `json:"lastname"`
	Firstname 	 string `json:"firstname"`
	Email 		 string `gorm:"unique" json:"email"`
    Phone        string `gorm:"unique" json:"phone"`
    Nationality  string `json:"nationality"`
}
