package jwts

import (
	"errors"
	"fast_gin/global"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type Claims struct {
	UserID uint `json:"user_id"`
	RoleID int8 `json:"role_id"`
}

type MyClaims struct {
	Claims
	jwt.RegisteredClaims
}

// SetToken 生成token
func SetToken(data Claims) (string, error) {
	SetClaims := MyClaims{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), //有效时间
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	// 使用指定的加密方式和声明类型创建新令牌
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	// 获得完整的签名令牌
	token, err := tokenStruct.SignedString([]byte(global.Config.Jwt.Key))
	if err != nil {
		logrus.Errorf("颁发jwt失败：%s", err)
		return "", err
	}
	return token, nil
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	// 解析、验证并返回token
	tokenObj, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Key), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tokenObj.Claims.(*MyClaims)
	if ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token无效")
	}

}
