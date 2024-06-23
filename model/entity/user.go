package entity

import "gorm.io/gorm"

/**
 * @title: logger
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64;not null"`
	Password string `json:"password" gorm:"size:128"`
	Image    string `json:"image" gorm:"size:255"`
}
