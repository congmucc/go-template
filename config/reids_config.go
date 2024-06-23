package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	error2 "gotemplate/utils/error"
	"time"
)

/**
 * @title: reids
 * @description:
 * @author: congmu
 * @date:    2024/6/23 15:02
 * @version: 1.0
 */

var redisClient *redis.Client

var redisDuration = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func InitRedis() *RedisClient {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     GlobalConfig.Redis.Url,
		Password: GlobalConfig.Redis.Password,
		DB:       GlobalConfig.Redis.DB,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		initErr = error2.AppendError(initErr, err)
		zLogger.Errorf("Load Redis Error: %s", err.Error())
	}

	return &RedisClient{}
}

func (rc *RedisClient) Set(key string, value any) error {
	return redisClient.Set(context.Background(), key, value, redisDuration).Err()
}
func (rc *RedisClient) Get(key string) (any, error) {
	return redisClient.Get(context.Background(), key).Result()
}
func (rc *RedisClient) Del(key string) error {
	return redisClient.Del(context.Background(), key).Err()
}
