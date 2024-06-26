package conf

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
 * @title: init
 * @description:
 * @author: congmu
 * @date:    2024/6/23 10:27
 * @version: 1.0
 */

var (
	// 异常处理，
	initErr error

	// GlobalConfig 存储配置文件全局变量
	GlobalConfig Config

	// GlobalLogger 暴露出来的日志对象
	GlobalLogger *zap.SugaredLogger

	// DB 暴露出来的数据库对象
	DB *gorm.DB

	// RedisTemplate 暴露出来的redis对象
	RedisTemplate *RedisClient
)

// 一般来说需要新建一个文件来存储全局变量，按理来说不需要这个的，但是为了标准写了
var zLogger = GlobalLogger

func init() {
	GlobalConfig = InitConfig()
	GlobalLogger = InitLogger()
	DB = InitDB()
	RedisTemplate = InitRedis()
}
