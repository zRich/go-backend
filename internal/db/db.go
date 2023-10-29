package db

import (
	"fmt"

	"github.com/zRich/go-backend/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Database interface {
	Connect() error
	SyncDatabase()
}

type Postgres struct {
	Config *DBConfig
}

//

func NewDatabase(config *DBConfig) Database {
	return &Postgres{
		Config: config,
	}
}

func (p *Postgres) SyncDatabase() {
	//自动迁移
	DB.AutoMigrate(&models.User{})
}

func (p *Postgres) Connect() error {
	//根据配置生成 postgresql 的连接
	config := p.Config
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("fatal error connect database: %w", err))
	}
	return nil
}
