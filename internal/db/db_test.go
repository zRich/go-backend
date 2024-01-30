package db

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitDB(t *testing.T) {
	viper.AddConfigPath("../../config/")
	viper.ReadInConfig()
	// config, err := GlobalConfig()
	// if err != nil {
	// 	t.Error(err)
	// }
	// _, err = InitDB(config)
	// if err != nil {
	// 	t.Error(err)
	// }
	database, err := InitDBFromConfig()
	database.AutoMigrate()
	if err != nil {
		t.Error(err)
	}

}
