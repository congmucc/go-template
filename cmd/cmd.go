package cmd

import (
	"gotemplate/config"
	"gotemplate/model/entity"
	"gotemplate/routers"
	"log"
)

/**
 * @title: logger
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

func Start() {
	// 自动迁移
	config.AutoMigrateDB(&entity.User{})
	// 路由初始化，这里会被阻塞，需要放在最后
	routers.InitRouter()
}

func Clear() {
	log.Println("Server exiting")
}
