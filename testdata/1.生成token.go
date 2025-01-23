package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/utils/jwts"
	"fmt"
)

func main() {

	core.InitLoggers()
	flags.Parse()
	global.Config = core.ReadConfig()
	token, err := jwts.SetToken(jwts.Claims{UserID: 1, RoleID: 1})

	fmt.Println(token, err)

	claims, err := jwts.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJpc3MiOiJjaGVueGkiLCJleHAiOjE3Mzc2NjEyMTN9.444wcDQ8w5AUllMmhMoNS4boQL4kIXejtWXsXC2KR4Y")
	fmt.Println(claims, err)
}
