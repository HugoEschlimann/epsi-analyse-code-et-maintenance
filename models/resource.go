package models

type ResourceType string

const (
	Book  ResourceType = "book"
	Game  ResourceType = "game"
	Film  ResourceType = "film"
	Autre ResourceType = "autre"
)

type Resource struct {
	ID 			uint `gorm:"primaryKey" json:"id"`
	Title 		string `json:"title"`
	Type 		ResourceType `json:"type" gorm:"type:enum('book', 'game', 'film', 'autre');default:'autre'"`
	Author 		string `json:"author"`
	IsAvailable bool `gorm:"default:true" json:"is_available"`
}