package config

type JwtConfig struct {
	Expires int    `yaml:"expires"`
	Issuer  string `yaml:"issuer"`
	Key     string `yaml:"key"`
}
