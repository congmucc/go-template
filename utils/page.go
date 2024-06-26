package utils

import "gorm.io/gorm"

/**
 * @title: page_dto
 * @description: 关于分页的通用类，详情请看对用户的查看。
 * @author: congmu
 * @date:    2024/6/26 14:53
 * @version: 1.0
 */

type PageRequest struct {
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

type Page[T any] struct {
	CurrentPage int64
	PageSize    int64
	Total       int64
	Data        []T
}

func (page *PageRequest) GetPageSize() int {
	if page.PageSize <= 0 {
		page.PageSize = 10
	}
	if page.PageSize >= 100 {
		page.PageSize = 100
	}
	return page.PageSize
}

func (page *PageRequest) GetPageNum() int {
	if page.PageNum <= 0 {
		page.PageNum = 1
	}
	return page.PageNum
}

// Paginate 分页函数处理
func Pageinate(page PageRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page.GetPageNum() - 1) * page.GetPageSize()).Limit(page.GetPageSize())
	}
}
