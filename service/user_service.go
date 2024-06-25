package service

import (
	"context"
	"gotemplate/conf"
	"gotemplate/dao"
	"gotemplate/model/dto"
	"gotemplate/model/entity"
)

/**
 * @title: user_service
 * @description:
 * @author: congmu
 * @date:    2024/6/25 15:43
 * @version: 1.0
 */

var zLogger = conf.GlobalLogger

type IUserService interface {
	UserLogin(ctx context.Context, userDto dto.UserDto)
}

// userService 结构体实现UserService接口
type UserService struct {
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		Dao: dao.NewUserDao(),
	}
}

func (u *UserService) UserLogin(ctx context.Context, userDto dto.UserDto) (entity.User, error) {
	// 盐加密，这里只做简单的盐
	userDto.Password = userDto.Password + "salt"
	login, err := u.Dao.GetUserByNameAndPassword(ctx, userDto)
	if err != nil {
		zLogger.Errorf("user login error: %v", err)
	}
	return login, err
}
