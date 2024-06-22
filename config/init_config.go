package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

/**
 * @title: logger
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

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

type LoggerConfig struct {
	MaxAge     int
	MaxSize    int
	MaxBackups int
	Level      string
}

type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
	Logger LoggerConfig
}

// GlobalConfig 存储配置文件全局变量
var GlobalConfig Config

func init() {
	// 设置配置文件的名字
	viper.SetConfigName("application")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")

	// 获取当前文件的绝对路径
	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get the current file path")
	}
	// 根据当前文件路径获取其所在的目录
	configPath := filepath.Dir(currentFilePath)

	// 添加配置文件的路径
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprint("Error reading config file: %s", err.Error()))
	}
	// 将配置文件内容解析到结构体中
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprint("Unable to decode into struct: %s", err.Error()))
	}
	GlobalConfig = config
}
