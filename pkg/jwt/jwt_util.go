package jwt

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v4"
)

// getExpireTime 过期时间
func getExpireTime() *jwtpkg.NumericDate {
	timenow := app.TimenowInTimezone()

	var expireTime int
	if config.Get[bool]("app.debug") {
		expireTime = config.Get[int]("jwt.debug_expire_time")
	} else {
		expireTime = config.Get[int]("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return jwtpkg.NewNumericDate(timenow.Add(expire))
}

// getTokenFromHeader 使用 jwtpkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func getTokenFromHeader(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		return "", ErrHeaderEmpty
	}
	if len(token) < 7 {
		return "", ErrHeaderMalformed
	}
	return token[7:], nil
}
