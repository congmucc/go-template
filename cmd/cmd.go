package cmd

import (
	"gotemplate/routers"
	"log"
)

/**
 * @title: logger
 * @description:
 * @author: congmu
 * @date:    2024/6/22 20:01
 * @version: 1.0
 */

func Start() {
	routers.InitRouter()
}

func Clear() {
	log.Println("Server exiting")
}
