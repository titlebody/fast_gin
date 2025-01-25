package config

type Upload struct {
	MaxSize  int64    `yaml:"max_size"`
	Dir      string   `yaml:"dir"`
	AllowExt []string `yaml:"allow_ext"`
}
