package dto

/**
 * @title: loginDto
 * @description:
 * @author: congmu
 * @date:    2024/6/22 19:15
 * @version: 1.0
 */

type UserDto struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
