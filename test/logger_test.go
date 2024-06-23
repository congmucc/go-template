package test

import (
	"gotemplate/config"
	"testing"
)

/**
 * @title: logger_test
 * @description:
 * @author: congmu
 * @date:    2024/6/22 21:58
 * @version: 1.0
 */

var zLogger = config.GlobalLogger

// 这只是测试如何使用，以及输出日志结构，这个log包整个可以删除
func TestInitLogger(t *testing.T) {
	zLogger.Error("出错了")
}
