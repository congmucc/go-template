package conf

import "gotemplate/api/controller"

/**
 * @title: init_controller
 * @description:
 * @author: congmu
 * @date:    2024/6/25 23:01
 * @version: 1.0
 */

func InitController() {
	controller.NewUserController()
}
