package entity

/**
 * @title: logger
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

type user struct {
	id       string `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
	image    string `json:"image"`
}
