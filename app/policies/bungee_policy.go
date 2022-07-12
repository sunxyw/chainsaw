package policies

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

// func CanModifyBungee(c *gin.Context, bungeeModel bungee.Bungee) bool {
//     return auth.CurrentUID(c) == bungeeModel.UserID
// }

func CanViewBungee(c *gin.Context) bool {
	return true
	return lo.Contains(c.GetStringSlice("perm_scopes"), "bungee:view")
}

// func CanCreateBungee(c *gin.Context, bungeeModel bungee.Bungee) bool {}
// func CanUpdateBungee(c *gin.Context, bungeeModel bungee.Bungee) bool {}
// func CanDeleteBungee(c *gin.Context, bungeeModel bungee.Bungee) bool {}
