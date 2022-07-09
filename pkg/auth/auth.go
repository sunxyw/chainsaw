// Package auth 授权相关逻辑
package auth

import "github.com/gin-gonic/gin"

// Attempt 尝试登录
// func Attempt(identifier string, password string) (user.User, error) {
// 	authRecord := user_auth.GetByIdentifier(identifier)
// 	if authRecord.ID == 0 {
// 		return user.User{}, errors.New("账号不存在")
// 	}

// 	if !authRecord.ComparePassword(password) {
// 		return user.User{}, errors.New("密码错误")
// 	}

// 	return user.Get(cast.ToString(authRecord.UserID)), nil
// }

// CurrentUser 从 gin.context 中获取当前登录用户
// func CurrentUser(c *gin.Context) user.User {
// 	userModel, ok := c.MustGet("current_user").(user.User)
// 	if !ok {
// 		logger.LogIf(errors.New("无法获取用户"))
// 		return user.User{}
// 	}
// 	// db is now a *DB value
// 	return userModel
// }

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
