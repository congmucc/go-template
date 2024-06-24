package error

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

/**
 * @title: validate_error
 * @description: 验证器返回指定错误信息
 * @author: congmu
 * @date:    2024/6/24 13:59
 * @version: 1.0
 */

// 解析验证器返回的错误信息
func ParseValidateError(errs validator.ValidationErrors, target interface{}) error {
	var errResult error
	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errs {
		// 获取字段名称
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintln("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}

		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}

// 使用如下，以login为例：
// 1. 在router中定义注册自定义验证器
// 参考router.registryCustValidator()方法。

// 2. 首先需要先在绑定的结构体内定义自定义错误：

// type UserDto struct {
//	Username string `form:"username" json:"username" binding:"required,bookabledate" message:"用户名错误" required_err:"用户名不能为空"`
//	Password string `form:"password" json:"password" binding:"required" message:"密码错误"`
//}

// 3. 其次需要在login调用

//func (userController UserController) Login(ctx *gin.Context) {
//	var userDto dto.UserDto
//	if err := ctx.ShouldBindJSON(&userDto); err != nil {
//		// 这里面可以做一个简易的自定义错误链
//		ctx.AbortWithStatusJSON(http.StatusBadRequest, result.ERR.ErrorWithMessage(error.ParseValidateError(err.(validator.ValidationErrors), &userDto).Error()))
//		return
//	}
//	ctx.JSON(http.StatusOK, result.OK.SuccessWithData(userDto))
//}

// 结果为：

//{
//    "username": "4",
//    "password": "123456"
//}

//{
//    "code": 1,
//    "message": "用户名错误"
//}
