package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//user's unique id,
	Email string `gorm:"uniqueIndex" json:"email,omitempty"`
	//user's credential
	Password string `json:"password,omitempty"`
	//user's role, 0 for admin, 1 non-admin, default is 1
	Role int `gorm:"default:1" json:"role,omitempty"`
	//user status, 0 for active, 1 for inactive, default is 0
	Status int `gorm:"default:0" json:"status,omitempty"`
}
