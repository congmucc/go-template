package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		//rgPublic.POST("/login", Login)
		rgPublic.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": map[string]any{
					"msg": "Login Success",
				},
			})
		})

		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "username": "zs"},
				},
			})
		})
	})
}
