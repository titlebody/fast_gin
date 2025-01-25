package routers

import (
	"fast_gin/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func Run() {
	gin.SetMode(global.Config.System.Mode)
	r := gin.Default()
	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")

	api := r.Group("/api")
	routerGroup := RouterGroup{api}

	routerGroup.UserRouter()
	routerGroup.ImageRouter()

	addr := global.Config.System.Addr()
	if global.Config.System.Mode == "release" {
		logrus.Infof("后端服务已启动 %s", addr)
	}

	r.Run(addr)

}
