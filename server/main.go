package main

import (
	"github.com/w871507855/BPanel/server/global"
	"github.com/w871507855/BPanel/server/initialize"
)

func main() {
	r := initialize.Routers()
	initialize.Viper()
	global.DB = initialize.InitDB()
	panic(r.Run(global.CONF.System.Addr))
}
