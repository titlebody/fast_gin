package captcha_api

import (
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
	"image/color"
)

type GenerateResponse struct {
	CaptchaID string `json:"captcha_id"`
	Base64    string `json:"base64"`
}

var digitDriver = &base64Captcha.DriverString{
	Height:          32,                                        //验证码高度
	Width:           100,                                       //验证码宽度
	NoiseCount:      0,                                         //验证码的噪声点
	ShowLineOptions: 2,                                         //显示是否有分割线
	Length:          4,                                         //生成验证码的长度
	Source:          "abcdefghijklmnopqrstuvwxyz",              //验证码中包含的字符合集
	BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125}, //验证码背景颜色
	Fonts:           []string{"wqy-microhei.ttc"},              //验证码使用的字体
}

func (CaptchaApi) GenerateView(c *gin.Context) {
	driver := digitDriver.ConvertFonts()                //为了处理字体
	b := base64Captcha.NewCaptcha(driver, CaptchaStore) //使用NewCaptcha这个方法接收两个参数来生成验证码
	id, b64s, _, err := b.Generate()                    //调用Generate方法得到此验证码的id，base64编码，验证码本身
	if err != nil {
		logrus.Errorf("图片验证码生成失败：%s", err)
		res.FailWithMsg("验证码生成失败", c)
		return
	}
	res.OKWithData(GenerateResponse{
		CaptchaID: id,
		Base64:    b64s,
	}, c)
}
