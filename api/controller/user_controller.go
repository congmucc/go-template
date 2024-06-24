package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gotemplate/model/dto"
	"gotemplate/utils/error"
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
}

func GetUserController() UserController {
	return UserController{}
}

func (userController UserController) Login(ctx *gin.Context) {
	var userDto dto.UserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		// 这里面可以做一个简易的自定义错误链
		ctx.AbortWithStatusJSON(http.StatusBadRequest, result.ERR.ErrorWithMessage(error.ParseValidateError(err.(validator.ValidationErrors), &userDto).Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.OK.SuccessWithData(userDto))
}
