package routers

import (
	"fast_gin/api"
	"fast_gin/middleware"
)

func (g RouterGroup) ImageRouter() {
	app := api.App.ImageApi
	g.POST("image/upload", middleware.AuthMiddleware, app.UploadView)
}
