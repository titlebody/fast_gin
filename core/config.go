package core

import (
	"fast_gin/config"
	"fast_gin/flags"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

// ReadConfig 读取配置文件
func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	file, err := os.ReadFile(flags.Options.File)
	if err != nil {
		logrus.Fatalf("配置文件读取失败:%s", err)
		return
	}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		logrus.Fatalf("配置文件格式错误：:%s", err)
		return
	}
	logrus.Infof("%s配置文件读取成功", flags.Options.File)
	return
}

// DumpConfig 写入配置文件
func DumpConfig(cfg *config.Config) {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		logrus.Errorf("配置文件转换错误:%s", err)
		return
	}
	err = os.WriteFile(flags.Options.File, data, 0644)
	if err != nil {
		logrus.Errorf("配置文件写入错误:%s", err)
		return
	}
	logrus.Infof("%s配置文件写入成功", flags.Options.File)
}
