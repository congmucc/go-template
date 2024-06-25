package controller

import (
	"github.com/gin-gonic/gin"
	"gotemplate/model/dto"
	"gotemplate/service"
	"gotemplate/utils/result"
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
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func GetUserController() UserController {
	return UserController{}
}

func (user UserController) Login(ctx *gin.Context) {
	var userDto dto.UserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		// 这里面可以做一个简易的自定义错误链
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result.ERR.ErrorWithMessage("传输信息错误"))
		return
	}

	login, err := user.UserService.UserLogin(ctx, userDto)
	if err != nil {
		ctx.JSON(http.StatusOK, result.ERR.ErrorWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.OK.SuccessWithData(login))
}
