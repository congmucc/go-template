package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gotemplate/utils"
	"log"
	"os"
	"time"
)

/**
 * @title: init_db
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

var (
	DB      *gorm.DB
	zLogger = GlobalLogger
)

// 不需要手动关闭，因为gorm会自动关闭
func InitDB() {
	logLevel := logger.Info
	if viper.GetString("profiles.active") == "dev" {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logLevel,    // Log level
			Colorful:      true,        // 彩色打印
		},
	)

	var err error
	dsn := ToDSN(GlobalConfig.Mysql)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", // 表名前缀，`User`表为 `sys_user`
			SingularTable: true,   // 禁用表名复数
		},
	})
	if err != nil {
		utils.AppendError(nil, err)
		zLogger.Errorf("Load Mysql configation Error: %s", err.Error())
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(GlobalConfig.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(GlobalConfig.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// 数据库迁移
func AutoMigrateDB(dst interface{}) {
	// 迁移 schema
	err := DB.AutoMigrate(dst)
	if err != nil {
		zLogger.Errorf("AutoMigrateDB Error: %s", err.Error())
	}
}

// ToDSN 将MysqlConfig转换为DSN字符串
func ToDSN(cfg MysqlConfig) string {
	//想要正确的处理time.Time,需要带上 parseTime 参数，
	//要支持完整的UTF-8编码，需要将 charset=utf8 更改为 charset=utf8mb4
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}
