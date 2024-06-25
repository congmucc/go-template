package dao

import (
	"context"
	"gotemplate/model/dto"
	"gotemplate/model/entity"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	return &UserDao{
		BaseDao: NewBaseDao(),
	}
}

func (u *UserDao) GetUserByNameAndPassword(ctx context.Context, userDto dto.UserDto) (entity.User, error) {
	var iUser entity.User
	err := u.Where("username = ? and password = ?", userDto.Username, userDto.Password).Find(&iUser)
	return iUser, err.Error
}
