// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
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

func IssueToken(uid string, ttype TokenType) string {
	uid = ttype.Prefix() + uid
	return JWT.Provider.IssueToken(uid)
}

func ParseToken(token string, ttype TokenType) (id string, err error) {
	id, err = JWT.Provider.ParseToken(token)
	if err != nil {
		return "", err
	}

	// 检查令牌前缀是否为对应类型
	if lo.Substring(id, 0, 4) != ttype.Prefix() {
		return "", ErrTokenInvalid
	}

	// 约定令牌前缀为4字符长，例 usr_, svc_
	id = id[4:]

	return id, nil
}

func ParseHeaderToken(c *gin.Context, ttype TokenType) (id string, err error) {
	token, err := getTokenFromHeader(c)
	if err != nil {
		return "", err
	}

	id, err = ParseToken(token, ttype)
	if err != nil {
		return "", err
	}

	return id, nil
}
