package user_api

import "github.com/gin-gonic/gin"

func (UserApi) UserListView(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})

}
