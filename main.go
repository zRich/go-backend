/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/zRich/go-backend/api"
	"github.com/zRich/go-backend/internal/db"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	// // 从配置文件中读取 postgresql 的配置
	// dbConfig, err := db.GlobalConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w ", err))
	// }
	// // 初始化数据库连接
	database, err := db.InitDBFromConfig()
	database.AutoMigrate()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	restConfig := api.RestServerConfig{}

	restConfig.Address = viper.GetString("server.address")
	restConfig.Port = viper.GetInt("server.port")
	restConfig.Prefix = viper.GetString("server.prefix")
	restConfig.DB = database

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	// restServer := server.NewServer(restConfig, database)
	restServer := api.NewRestServer(restConfig)

	err = restServer.Start()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

}
