package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string
}

type MysqlConfig struct {
	Url      string
	Username string
	Password string
	DBName   string
	Port     string
}

type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
}

func InitConfig() Config {
	// 设置配置文件的名字
	viper.SetConfigName("application")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprint("Error reading config file: %s", err.Error()))
	}
	// 将配置文件内容解析到结构体中
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprint("Unable to decode into struct: %s", err.Error()))
	}
	return config
}
