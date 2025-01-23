package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/routers"
)

func main() {
	core.InitLoggers()
	flags.Parse()
	global.Config = core.ReadConfig()
	core.InitGorm()
	core.InitRedis()
	flags.Run()

	routers.Run()

}
