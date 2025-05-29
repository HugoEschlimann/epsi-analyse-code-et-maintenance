package models

import "gorm.io/gorm"

type ResourceType string

const (
	Book ResourceType = "book"
	Game ResourceType = "game"
	Film ResourceType = "film"
)

type Resource struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Title string
	Type ResourceType
	Author string
	IsAvailable bool
}