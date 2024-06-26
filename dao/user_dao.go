package dao

import (
	"context"
	"gotemplate/model/dto"
	"gotemplate/model/entity"
	"gotemplate/utils"
)

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	return &UserDao{
		BaseDao: NewBaseDao(),
	}
}

// 获取用户名和密码
func (u *UserDao) GetUserByNameAndPassword(ctx context.Context, userDto dto.UserDto) (entity.User, error) {
	var iUser entity.User
	err := u.Where("username = ? and password = ?", userDto.Username, userDto.Password).Find(&iUser)
	return iUser, err.Error
}

// 分页查询
func (u *UserDao) GetPage(ctx context.Context, pageDto utils.PageRequest) (utils.Page[dto.UserDto], error) {
	var userList utils.Page[dto.UserDto]
	u.Model(&entity.User{}).
		Scopes(utils.Pageinate(pageDto)).
		Find(&userList.Data).
		Offset(-1).Limit(-1).Count(&userList.Total)
	return userList, nil

}
