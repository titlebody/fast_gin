package user_api

import (
	"fast_gin/middleware"
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
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

	fmt.Println(cr)
	res.OKWithData("用户登录", c)
}
