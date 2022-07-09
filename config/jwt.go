package config

import "gohub/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{

			// 令牌提供者，支持：local、firebase
			// 请注意，由 firebase 签发的令牌无法直接使用，需要由客户端转换为 ID Token
			// 建议你仅在需要与 firebase 相关的功能时使用 firebase 提供者
			// 如在前端使用 firebase realtime database 等
			"token_provider": config.Env("JWT_TOKEN_PROVIDER", "local"),

			// 过期时间，单位是分钟，一般不超过两个小时
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120),

			// 允许刷新时间，单位分钟，86400 为两个月，从 Token 的签名时间算起
			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 86400),

			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_time": 86400,

			// 启用黑名单，允许从服务端登出用户
			"enable_blacklist": config.Env("JWT_ENABLE_BLACKLIST", false),
		}
	})
}
