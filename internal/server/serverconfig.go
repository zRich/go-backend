package server

import "github.com/spf13/viper"

type RestConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// 从配置文件中读取配置，返回 config 结构体

func GlobalConfig() (*RestConfig, error) {
	config := &RestConfig{}
	if err := config.load(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *RestConfig) load() error {
	serverAddress := viper.GetString("server.address")
	serverPort := viper.GetInt("server.port")
	c.Address = serverAddress
	c.Port = serverPort

	return nil
}
