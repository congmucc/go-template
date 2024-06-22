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

// 存储配置文件全局变量，防止频繁读取配置文件
var GlobalConfig Config

func InitConfig() {
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
	/**
	 * 这里还可以使用key value读取配置文件
	 * viper.GetString("server.port")
	 */
	if err != nil {
		panic(fmt.Sprint("Unable to decode into struct: %s", err.Error()))
	}
	GlobalConfig = config
}
