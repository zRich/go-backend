package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// useer name, default is ""
	Name  string `gorm:"default:''"`
	Email *string
	//student's password in hash，default is ""
	Password string `gorm:"default:''"`
}
