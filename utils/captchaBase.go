package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 设置自带的 store（可以配置成redis）
var store = base64Captcha.DefaultMemStore

type CaptchaResponse struct {
	Id     string
	Encode string
	Answer string
}

func Captcha() *CaptchaResponse {
	// 配置验证码的参数
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	// ConvertFonts 按名称加载字体
	driver := driverString.ConvertFonts()
	// 创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	// Generate 生成随机 id、base64 图像字符串
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return nil
	}
	return &CaptchaResponse{
		Id:     id,
		Encode: b64s,
		Answer: answer,
	}
}
