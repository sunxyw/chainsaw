// Package config 站点配置信息
package config

import "gohub/pkg/config"

func init() {
	config.Add("auth", func() map[string]interface{} {
		return map[string]interface{}{

			// 注册相关配置
			"register": map[string]interface{}{
				// 是否开启注册功能
				"enable": config.Env("AUTH_REGISTER_ENABLE", true),

				// 注册是否需要验证码
				"captcha": config.Env("AUTH_REGISTER_CAPTCHA", false),
			},

			// 登录相关配置
			"login": map[string]interface{}{
				// 是否开启登录功能
				"enable": config.Env("AUTH_LOGIN_ENABLE", true),

				// 登录是否需要验证码
				"captcha": config.Env("AUTH_LOGIN_CAPTCHA", true),
			},

			// 允许的登录类型
			"login_types": config.Env("AUTH_LOGIN_TYPE", []string{"email", "phone", "username"}),
		}
	})
}
