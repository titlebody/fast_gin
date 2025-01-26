package api

import (
	"fast_gin/api/captcha_api"
	"fast_gin/api/image_api"
	"fast_gin/api/user_api"
)

type Api struct {
	UserApi  user_api.UserApi
	ImageApi image_api.ImageApi
	Captcha  captcha_api.CaptchaApi
}

var App = new(Api)
