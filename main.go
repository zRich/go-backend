/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/server"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	// 从配置文件中读取 postgresql 的配置
	dbConfig := &db.DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
	}

	db := db.NewDatabase(dbConfig)
	db.Connect()
	db.SyncDatabase()

	port := viper.GetInt("server.port")
	fmt.Println(port)
	httpConfig := &server.ServerConfig{
		Port: port,
	}
	server := server.NewServer(httpConfig)
	err = server.Start()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
}
