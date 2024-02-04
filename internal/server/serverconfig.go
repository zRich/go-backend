package server

type ServerConfig interface {
	GetAddress() string
	GetPort() int
	GetVersion() string
}

// type RestConfig struct {
// 	Address string `json:"address"`
// 	Port    int    `json:"port"`
// }

// func (c *RestConfig) GetAddress() string {
// 	return c.Address
// }

// func (c *RestConfig) GetPort() int {
// 	return c.Port
// }

// // 从配置文件中读取配置，返回 config 结构体

// func GlobalConfig() (ServerConfig, error) {
// 	config := &RestConfig{}
// 	if err := config.load(); err != nil {
// 		return nil, err
// 	}

// 	return config, nil
// }

// func (c *RestConfig) load() error {
// 	serverAddress := viper.GetString("server.address")
// 	serverPort := viper.GetInt("server.port")
// 	c.Address = serverAddress
// 	c.Port = serverPort

// 	return nil
// }
