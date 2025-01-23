package user_api

import "github.com/gin-gonic/gin"

func (UserApi) LoginView(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
