package core

import (
	"fast_gin/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func InitGorm() {
	var cfg = global.Config
	Dialectic := cfg.DB.DSN()
	if Dialectic == nil {
		return
	}
	db, err := gorm.Open(Dialectic, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //不生成实体外键
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败：%s", err)
	}
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("数据库连接失败：%s", err)
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		logrus.Fatalf("数据库连接失败：%s", err)
		return
	}
	// 设置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功！")
	global.DB = db
	return
}
