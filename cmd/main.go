package main

import (
	"gotemplate/conf"
	"gotemplate/model/entity"
	"gotemplate/routers"
)

func main() {
	// 自动迁移
	conf.AutoMigrateDB(&entity.User{})

	// 路由初始化，这里会被阻塞，需要放在最后
	routers.InitRouter()
}
