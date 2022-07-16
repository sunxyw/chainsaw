// Package routes 注册路由
package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/middlewares"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get[string]("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.ThrottleByIP("200-H"), middlewares.Perm())
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.ThrottleByIP("1000-H"))
		{
		}

		bungeeGroup := v1.Group("/bungees")
		{
			bgc := new(controllers.BungeesController)
			bungeeGroup.GET("", bgc.Index)
		}

		notificationGroup := v1.Group("/notifications")
		{
			ntc := new(controllers.NotificationsController)
			notificationGroup.GET("", ntc.Index)
			notificationGroup.GET("/next", ntc.GetNextQueued)
			notificationGroup.GET("/:id", ntc.Show)
		}

		newsGroup := v1.Group("/news")
		{
			nsc := new(controllers.NewsController)
			newsGroup.GET("", nsc.Index)
		}

		mapsGroup := v1.Group("/maps")
		{
			mcc := new(controllers.McmapsController)
			mapsGroup.GET("", mcc.Index)
			mapsGroup.GET("/:id", mcc.Show)
			mapsGroup.POST("", mcc.Store)
		}
	}
}
