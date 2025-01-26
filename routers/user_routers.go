package routers

import (
	"fast_gin/api"
	"fast_gin/api/user_api"
	"fast_gin/middleware"
	"fast_gin/models"
)

func (g RouterGroup) UserRouter() {
	app := api.App.UserApi
	g.POST("users/login", middleware.LimitMiddleware(3), middleware.BindJsonMiddleware[user_api.LoginRequest], app.LoginView)
	g.POST("users/logout", middleware.AuthMiddleware, app.LogoutView)
	g.GET("users", middleware.LimitMiddleware(10), middleware.AdminMiddleware, middleware.BindQueryMiddleware[models.PageInfo], app.UserListView)

}
