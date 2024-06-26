package routers

import (
	"github.com/gin-gonic/gin"
	"gotemplate/api/controller"
	"net/http"
)

/**
 * @title: logger
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userGroup := rgPublic.Group("user")
		{
			userGroup.POST("/login", controller.NewUserController().Login)
		}

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
