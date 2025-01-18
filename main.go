package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitLoggers()
	flags.Run()
	global.Config = core.ReadConfig()
	
	logrus.Infof("你好")
	logrus.Warnf("你好")
	logrus.Error("你好")

}
