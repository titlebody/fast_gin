package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fmt"
)

func main() {
	core.InitLoggers()
	flags.Parse()
	global.Config = core.ReadConfig()
	core.InitGorm()
	core.InitRedis()
	flags.Run()

	fmt.Println("开启web服务")
}
