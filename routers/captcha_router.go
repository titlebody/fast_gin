package routers

import "fast_gin/api"

func (g RouterGroup) CaptchaRouter() {
	captchaApp := api.App.Captcha
	g.GET("captcha", captchaApp.GenerateView)

}
