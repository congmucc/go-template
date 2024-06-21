package cmd

import (
	"gotemplate/config"
	"gotemplate/routers"
)

func Start() {
	config.InitConfig()
	routers.InitRouter()
}

func Clear() {

}
