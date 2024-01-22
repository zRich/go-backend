package db

import (
	"fmt"

	"github.com/zRich/go-backend/internal/db/models"
	"github.com/zRich/go-backend/internal/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreDB struct {
	DB        *gorm.DB
	Config    *DBConfig
	connected bool
}

func NewPostgreDB(config *DBConfig) Database {
	postgres := &PostgreDB{}
	postgres.Config = config
	postgres.connected = false
	return postgres
}

func (p *PostgreDB) AutoMigrate() {
	// 自动迁移
	p.DB.AutoMigrate(&models.User{})
	p.DB.AutoMigrate(&models.Course{})
	p.DB.AutoMigrate(&models.Task{})
	p.DB.AutoMigrate(&models.Student{})
}

func (p *PostgreDB) Connect() (*gorm.DB, error) {

	if p.connected {
		log.Log.Info("database already connected")
		return p.DB, nil
	}

	config := p.Config
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
	fmt.Println(dsn)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("fatal error connect database: %w", err))
	}

	p.DB = DB
	p.connected = true
	return DB, nil
}
