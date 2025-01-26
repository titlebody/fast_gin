package redis_ser

import (
	"context"
	"fast_gin/global"
	"fast_gin/utils/jwts"
	"fmt"
	"time"
)

func Logout(token string) {
	claims, err := jwts.CheckToken(token)
	if err != nil {
		return
	}
	key := fmt.Sprintf("logout_%s", token)
	sub := claims.ExpiresAt.Sub(time.Now())
	_, _ = global.Redis.Set(context.Background(), key, "", sub).Result()
}

func HasLogout(token string) bool {
	key := fmt.Sprintf("logout_%s", token)
	_, err := global.Redis.Get(context.Background(), key).Result()
	if err == nil {
		return true
	}
	return false

}
