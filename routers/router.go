package routers

import (
	"context"
	"fmt"
	"gotemplate/config"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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

// 初始化路由，以及优雅退出
func InitRouter() {
	// 创建一个信号量，分别监听结束（ctrl+c）和退出应用这个信号，这里使用了context上下文
	// 参考：https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1/")

	InitBasePlatformRoutes()

	// 遍历router的数组，添加到路由中
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	// 对端口进行监听
	stPort := config.GlobalConfig.Server.Port
	if stPort == "" {
		panic(fmt.Sprint("Load Router Port Error"))
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// 在启动监听前先打印启动日志，并加入短暂延迟
	log.Printf("Starting server on port %s...\n", stPort)
	time.Sleep(100 * time.Millisecond) // 短暂延迟，给日志打印机会

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start server listen error: %s\n", err.Error())
		}
	}()

	// 监听信号，终止服务
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err.Error())
		return
	}
}

// 初始化基础模块路由
func InitBasePlatformRoutes() {
	InitUserRoutes()

}
