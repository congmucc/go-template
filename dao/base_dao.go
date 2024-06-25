package dao

import (
	"gorm.io/gorm"
	"gotemplate/conf"
)

/**
 * @title: base_dao
 * @description:
 * @author: congmu
 * @date:    2024/6/25 21:09
 * @version: 1.0
 */

type BaseDao struct {
	*gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		DB: conf.DB,
	}
}
