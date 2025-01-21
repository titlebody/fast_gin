package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
)

func main() {
	core.InitLoggers()
	flags.Run()
	global.Config = core.ReadConfig()
	core.InitGorm()
	core.InitRedis()
}
