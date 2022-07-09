package jwt

import (
	"gohub/pkg/cache"
	"gohub/pkg/config"
	"gohub/pkg/logger"

	jwtpkg "github.com/golang-jwt/jwt/v4"
)

type LocalProvider struct {
	SignKey []byte
}

type claims struct {
	UID string `json:"uid"`

	jwtpkg.RegisteredClaims
}

func NewLocalProvider() *LocalProvider {
	return &LocalProvider{
		SignKey: []byte(config.Get[string]("app.key")),
	}
}

// IssueToken 生成 Token
func (lt *LocalProvider) IssueToken(uid string) string {
	claims := claims{
		uid,
		jwtpkg.RegisteredClaims{
			ExpiresAt: getExpireTime(),
			Issuer:    config.Get[string]("app.name"),
		},
	}

	token, err := lt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

func (lt *LocalProvider) ParseToken(tokenString string) (uid string, err error) {
	token, err := lt.parseTokenString(tokenString)

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return "", ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return "", ErrTokenExpired
			}
		}
		return "", ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*claims); ok && token.Valid {

		// 检查 Token 是否已注销
		if lt.checkBlacklist(claims.ID) {
			return "", ErrTokenInvalid
		}

		return claims.UID, nil
	}

	return "", ErrTokenInvalid
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (lt *LocalProvider) createToken(claims claims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(lt.SignKey)
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (lt *LocalProvider) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &claims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return lt.SignKey, nil
	})
}

// checkBlacklist 检查 Token 是否在黑名单中
func (lt *LocalProvider) checkBlacklist(key string) bool {
	if ok := cache.Get("jwt:blacklist:" + key); ok != nil {
		return true
	}
	return false
}
