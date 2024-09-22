package routers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gotemplate/api/middleware"
	"gotemplate/conf"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

/**
 * @title: logger
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

// rgPublic: 不需要鉴权， rgAuth: 需要鉴权,携带token
type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 定义一个切片存储所有路由
var (
	gfnRoutes []IFnRegistRoute
	zLogger   = conf.GlobalLogger
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

	r.Use(middleware.Cors())

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1/")

	// 路由鉴权
	rgAuth.Use(middleware.Auth())

	// 初始化基础模块路由
	initBasePlatformRoutes()

	// 注册自定义验证器
	//registryCustValidator()

	// 遍历router的数组，添加到路由中
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	// 对端口进行监听
	stPort := conf.GlobalConfig.Server.Port
	if stPort == "" {
		panic(fmt.Sprint("Load Router Port Error"))
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// 在启动监听前先打印启动日志，并加入短暂延迟
	zLogger.Infof("Starting server on port %s", stPort)
	time.Sleep(100 * time.Millisecond) // 短暂延迟，给日志打印机会

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zLogger.Errorf("Start server listen error: %s", err.Error())
		}
	}()

	// 监听信号，终止服务
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zLogger.Errorf("Server shutdown error: %s", err.Error())
		return
	}
}

// 初始化基础模块路由
func initBasePlatformRoutes() {
	InitUserRoutes()

}

// 注册自定义验证器
// 参考：https://gin-gonic.com/zh-cn/docs/examples/custom-validators/
// 在本框架中请在utils/error/validate_error.go查看相应例子
//func registryCustValidator() {
//	// 由于这里不适合所有人，先不写 下面这个例子是判断含有bookabledate这个校验属性的属性是否是以t开头，不然就返回错误
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		v.RegisterValidation("bookabledate", func(fl validator.FieldLevel) bool {
//			if value, ok := fl.Field().Interface().(string); ok {
//				if value != "" && 0 == strings.Index(value, "t") {
//					return true
//				}
//			}
//			return false
//		})
//	}
//}
