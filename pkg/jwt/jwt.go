// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"gohub/pkg/logger"
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

func ParseToken(token string) (id string, tokenType TokenType, err error) {
	id, err = JWT.Provider.ParseToken(token)

	if err != nil {
		return
	}

	if isServiceToken(id) {
		tokenType = TokenTypeService
		id = id[3:]
	} else {
		tokenType = TokenTypeUser
	}

	return
}

func ParseHeaderToken(c *gin.Context, tokenType TokenType) (id string, ttype TokenType, err error) {
	token, err := getTokenFromHeader(c)
	if err != nil {
		return
	}

	id, ttype, err = ParseToken(token)
	if err != nil {
		return "", ttype, err
	}

	inType := getTokenTypeList(tokenType)
	logger.Dump(inType)
	for _, t := range inType {
		if t == ttype {
			return id, ttype, nil
		}
	}

	return "", ttype, ErrTokenInvalid
}
