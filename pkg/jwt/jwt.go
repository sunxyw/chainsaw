// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWTService struct {
	Provider TokenProvider
}

var once sync.Once
var JWT *JWTService

func InitWithProvider(provider TokenProvider) {
	once.Do(func() {
		JWT = &JWTService{
			Provider: provider,
		}
	})
}

func IssueToken(uid string) string {
	return JWT.Provider.IssueToken(uid)
}

func ParseToken(token string) (uid string, err error) {
	return JWT.Provider.ParseToken(token)
}

func ParseHeaderToken(c *gin.Context) (uid string, err error) {
	token, err := getTokenFromHeader(c)
	if err != nil {
		return "", err
	}
	return ParseToken(token)
}
