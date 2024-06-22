package cmd

import (
	"gotemplate/config"
	"gotemplate/routers"
	"log"
)

func Start() {
	config.InitConfig()
	routers.InitRouter()
}

func Clear() {
	log.Println("Server exiting")
}
