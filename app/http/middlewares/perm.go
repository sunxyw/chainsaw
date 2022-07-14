package middlewares

import (
	"gohub/app/models/service_token"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func Perm() gin.HandlerFunc {
	return func(c *gin.Context) {

		var scopes []string

		id, err := jwt.ParseHeaderToken(c, jwt.TokenTypeService)
		if err != nil {
			scopes = []string{}
		} else {
			stModel := service_token.Get(id)
			if stModel.ID == 0 {
				response.Unauthorized(c, "找不到对应服务令牌，令牌可能已删除")
				return
			}

			scopes = strings.Split(stModel.Scopes, ",")
		}

		c.Set("perm_scopes", scopes)
		c.Next()
	}
}
