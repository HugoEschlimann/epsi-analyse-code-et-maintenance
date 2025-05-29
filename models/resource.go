package models

type ResourceType string

const (
	Book ResourceType = "book"
	Game ResourceType = "game"
	Film ResourceType = "film"
)

type Resource struct {
	ID uint `gorm:"primaryKey"`
	Title string
	Type ResourceType
	Author string
	IsAvailable bool
}