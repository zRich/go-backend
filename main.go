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

	// // 从配置文件中读取 postgresql 的配置
	// dbConfig, err := db.GlobalConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w ", err))
	// }
	// // 初始化数据库连接
	database, err := db.InitDBFromConfig()
	// err = db.InitDB(dbConfig)

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	restConfig, err := server.GlobalConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	restServer := server.NewServer(restConfig, *database)

	// 注册路由

	// 启动 http 服务
	err = restServer.Start()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

}
