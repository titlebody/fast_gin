package global

import (
	"fast_gin/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Version string = "0.0.1"
	Config  *config.Config
	DB      *gorm.DB
	Redis   *redis.Client
)
