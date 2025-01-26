package captcha_api

import "github.com/mojocn/base64Captcha"

type CaptchaApi struct {
}

var CaptchaStore = base64Captcha.DefaultMemStore //store不但存储id和对应验证码还包含了Get，Verify两个方法后续会用到
