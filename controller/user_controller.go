package controller

import (
	"github.com/gin-gonic/gin"
	"gotemplate/model/dto"
	"net/http"
)

/**
 * @title: userController
 * @description:
 * @author: congmu
 * @date:    2024/6/22 19:03
 * @version: 1.0
 */

type UserController struct {
}

func GetUserController() UserController {
	return UserController{}
}

func (userController UserController) Login(ctx *gin.Context) {
	var userDto dto.UserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]any{
			"username": userDto.Username,
			"password": userDto.Password,
			"msg":      "Login Success",
		},
	})
}
