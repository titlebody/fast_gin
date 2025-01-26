package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

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

var store = base64Captcha.DefaultMemStore //store不但存储id和对应验证码还包含了Get，Verify两个方法后续会用到

// 生成验证码
func CaptchaGenerate() (string, string, string, error) {
	driver := digitDriver.ConvertFonts()         //为了处理字体
	b := base64Captcha.NewCaptcha(driver, store) //使用NewCaptcha这个方法接收两个参数来生成验证码
	id, b64s, hcode, err := b.Generate()         //调用Generate方法得到此验证码的id，base64编码，验证码本身
	if err != nil {
		fmt.Println("Error generating captcha:", err)
		return "", "", "", err
	}
	fmt.Println("id:", id, "base64:", b64s, "hcode:", hcode)
	return id, b64s, hcode, nil
}
