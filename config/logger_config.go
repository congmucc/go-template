package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

/**
 * @title: logger_config
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

// 从配置文件中获取信息
var loggerConfig = GlobalConfig.ZLogger

// 初始化日志配置
func InitLogger() *zap.SugaredLogger {
	level := zapcore.InfoLevel
	if loggerConfig.Level == "debug" {
		level = zapcore.DebugLevel
	}
	//core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), level)
	core := zapcore.NewCore(getEncoder(), zapcore.AddSync(os.Stdout), level)
	return zap.New(core).Sugar()
}

// 对SugaredLogger配置初始化
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "msg"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 设置日志的输出属性
func getWriteSyncer() zapcore.WriteSyncer {
	logSeparator := string(filepath.Separator)
	logRootDir, _ := os.Getwd()
	// todo linux和windows存储日志文件如何在同一路径下，
	// 上面那个相对路径其实需要改一下，如果是测试类的话，相对路径容易出错
	logFilePath := logRootDir + logSeparator + "log" + logSeparator + time.Now().Format(time.DateOnly) + ".log"

	// 日志切割
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    loggerConfig.MaxSize,
		MaxAge:     loggerConfig.MaxAge,
		MaxBackups: loggerConfig.MaxBackups,
		LocalTime:  true,
		Compress:   loggerConfig.compress,
	}
	return zapcore.AddSync(lumberjackSyncer)
}
