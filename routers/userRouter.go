package routers

import (
	"github.com/gin-gonic/gin"
	"gotemplate/api"
	"net/http"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userGroup := rgPublic.Group("user")
		{
			userGroup.POST("/login", api.GetUserApi().Login)
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
