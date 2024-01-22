package models

import "gorm.io/gorm"

// Task is a model for a task in the database.
// a task belongs to a course, and a course has many tasks.

type Task struct {
	gorm.Model
	TaskNo      string  `gorm:"uniqueIndex" json:"taskNo,omitempty"`
	Description *string `json:"description,omitempty"`
	// This is a "belongs to" relationship. A task belongs to a course.
	// This is the "many" side of the relationship.
	CourseID uint `json:"courseID,omitempty"`
	//score of the task，default is 0
	Score int `gorm:"default:0" json:"score,omitempty"`
}
