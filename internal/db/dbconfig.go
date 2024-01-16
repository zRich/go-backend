package db

import "github.com/spf13/viper"

type DBConfig struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func GlobalConfig() (*DBConfig, error) {
	config := &DBConfig{}
	if err := config.load(); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *DBConfig) load() error {
	c.Type = viper.GetString("database.type")
	c.Host = viper.GetString("database.host")
	c.Port = viper.GetInt("database.port")
	c.User = viper.GetString("database.user")
	c.Password = viper.GetString("database.password")
	c.DBName = viper.GetString("database.dbname")

	return nil
}
