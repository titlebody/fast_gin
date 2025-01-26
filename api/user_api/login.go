package user_api

import (
	"fast_gin/api/captcha_api"
	"fast_gin/global"
	"fast_gin/middleware"
	"fast_gin/models"
	"fast_gin/utils/jwts"
	"fast_gin/utils/pwd"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

type LoginRequest struct {
	Username    string `json:"username" binding:"required" label:"用户名"`
	Password    string `json:"password" binding:"required" label:"密码"`
	CaptchaID   string `json:"captcha_id" label:"验证码ID"`
	CaptchaCode string `json:"captcha_code" label:"验证码"`
}

/*func ShowBindJson[T any](c *gin.Context) (cr T, err error) {
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	return
}*/

func (UserApi) LoginView(c *gin.Context) {

	cr := middleware.GetBind[LoginRequest](c)

	if global.Config.Site.Login.Captcha {
		if cr.CaptchaID == "" || cr.CaptchaCode == "" {
			res.FailWithMsg("请输入图片验证码", c)
			return
		}
		if !captcha_api.CaptchaStore.Verify(cr.CaptchaID, strings.ToLower(cr.CaptchaCode), true) {
			res.FailWithMsg("验证码错误", c)
			return
		}
	}

	var user models.UserModel

	err := global.DB.Take(&user, "username = ?", cr.Username).Error
	if err != nil {
		res.FailWithMsg("用户名或者密码错误", c)
		return
	}
	if err = pwd.ComparePasswords(user.Password, cr.Password); err != nil {
		res.FailWithMsg("用户名或者密码错误", c)
		return
	}
	token, err := jwts.SetToken(jwts.Claims{
		UserID: user.ID,
		RoleID: user.RoleID,
	})
	if err != nil {
		logrus.Errorf("生成Token失败：%s", err)
		res.FailWithMsg("登录失败", c)
		return
	}
	res.OKWithData(token, c)
}
