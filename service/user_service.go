package service

import (
	"context"
	"gotemplate/conf"
	"gotemplate/dao"
	"gotemplate/model/dto"
	"gotemplate/model/entity"
	"gotemplate/utils"
)

/**
 * @title: user_service
 * @description:
 * @author: congmu
 * @date:    2024/6/25 15:43
 * @version: 1.0
 */

var zLogger = conf.GlobalLogger

// userService 结构体实现UserService接口 参考go-micro这种类型的接口，这个方式可以学一下
//type UserService interface {
//	UserLogin(ctx context.Context, userDto dto.UserDto) (entity.User, error)
//	GetPage(ctx context.Context) (user []entity.User, err error)
//}

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

func (u *UserService) GetPage(c context.Context, pageDto utils.PageRequest) (utils.Page[dto.UserDto], error) {
	userList, err := u.Dao.GetPage(c, pageDto)
	if err != nil {
		zLogger.Errorf("user get page error: %v", err)
	}
	return userList, err
}
