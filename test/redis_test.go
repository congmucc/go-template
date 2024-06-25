package test

import (
	"fmt"
	"gotemplate/conf"
	"testing"
)

/**
 * @title: redis_test
 * @description:
 * @author: congmu
 * @date:    2024/6/23 18:43
 * @version: 1.0
 */

var redisTemplate = conf.RedisTemplate

func TestRedis(t *testing.T) {
	redisTemplate.Set("key", "你好")
	value, _ := redisTemplate.Get("key")
	fmt.Printf("key的值：%s\n", value)
}
