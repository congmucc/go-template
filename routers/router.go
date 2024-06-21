package routers

import (
	"fmt"
	"gotemplate/config"
)

import (
	"github.com/gin-gonic/gin"
) // rgPublic: 不需要鉴权， rgAuth: 需要鉴权,携带token

// rgPublic: 不需要鉴权， rgAuth: 需要鉴权,携带token
type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 定义一个切片存储所有路由
var (
	gfnRoutes []IFnRegistRoute
)

// 注册路由，向切片中添加路由
func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

// 初始化路由
func InitRouter() {
	r := gin.Default()

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1/")

	InitBasePlatformRoutes()

	// 遍历router的数组，添加到路由中
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	// 对端口进行监听
	stPort := config.InitConfig().Server.Port
	if stPort == "" {
		panic(fmt.Sprint("Load Router Port Error"))
	}

	err := r.Run(fmt.Sprint("%s", stPort))
	if err != nil {
		panic(fmt.Sprint("Start Server Error: %s", err.Error()))
	}

}

// 初始化基础模块路由
func InitBasePlatformRoutes() {
	InitUserRoutes()

}
