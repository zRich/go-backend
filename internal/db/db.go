package db

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database interface {
	Connect() (*gorm.DB, error)
}

func InitDB(config *DBConfig) (*Database, error) {
	var err error
	switch config.Type {
	case "postgres":
		database := NewPostgreDB(config)
		DB, err = database.Connect()
		if err != nil {
			return nil, err
		}
		return &database, nil
	default:
		panic("unsupported database type")
	}
	// return nil, nil
}

func InitDBFromConfig() (*Database, error) {
	config, err := GlobalConfig()
	if err != nil {
		return nil, err
	}
	return InitDB(config)
}
