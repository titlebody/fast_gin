package flags

import (
	"flag"
)

type FlagOptions struct {
	File string
}

var Options FlagOptions

func Run() {
	flag.StringVar(&Options.File, "f", "config.yaml", "配置文件路径")
	flag.Parse()
}
