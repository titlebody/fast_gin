package flags

import (
	"fast_gin/global"
	"fast_gin/models"
	"fast_gin/utils/pwd"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

type User struct {
}

func (User) Create() {
	var user models.UserModel
	fmt.Println("请选择角色: 1 管理员 2 普通用户")
	_, err := fmt.Scanln(&user.RoleID)
	if err != nil {
		fmt.Println("输入错误", err)
		return
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		fmt.Print("用户角色输入错误", err)
		return
	}
	fmt.Println("请输入用户名:")
	fmt.Scanln(&user.Username)
	// 判断用户名是否存在
	var u models.UserModel
	err = global.DB.Take(&u, "username = ?", user.Username).Error
	if err == nil {
		fmt.Println("用户名已存在！")
		return
	}
	fmt.Println("请输入密码:")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("请再次输入密码:")
	RePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if string(password) != string(RePassword) {
		fmt.Println("两次密码不一致")
		return
	}
	hashPwd, _ := pwd.HashPassword(string(password))
	err = global.DB.Create(&models.UserModel{
		Username: user.Username,
		RoleID:   user.RoleID,
		Password: hashPwd,
	}).Error
	if err != nil {
		logrus.Errorf("用户创建失败：%s", err)
		return
	}
	logrus.Infof("用户创建成功")

}

func (User) List() {
	var userList []models.UserModel
	err := global.DB.Order("created_at desc").Limit(10).Find(&userList).Error
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	for _, user := range userList {
		fmt.Printf("用户ID:%d 用户名：%s 用户昵称：%s 用户角色：%d 创建时间：%s \n", user.ID, user.Username, user.Nickname, user.RoleID, user.CreatedAt.Format("2006-01-02 15:04:05"))
	}

}
