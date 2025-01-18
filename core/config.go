package core

import (
	"fast_gin/config"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("配置文件读取失败:%s", err)
		return
	}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		fmt.Printf("配置文件格式错误：:%s", err)
		return
	}
	return
}
