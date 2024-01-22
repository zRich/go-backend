package models

import (
	"gorm.io/gorm"
)

// Course is a model for a course in the database.
type Course struct {
	gorm.Model
	Name        string  `gorm:"uniqueIndex" json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
