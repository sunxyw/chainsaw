package v1

import (
	"gohub/app/policies"
	"gohub/pkg/bungee"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type BungeesController struct {
	BaseAPIController
}

func (ctrl *BungeesController) Index(c *gin.Context) {

	if ok := policies.CanViewBungee(c); !ok {
		response.Forbidden(c)
		return
	}

	bungee.Cluster.Lock.RLock()
	defer bungee.Cluster.Lock.RUnlock()

	result := make(map[string]map[string][]bungee.BungeePlayer)

	for _, proxy := range bungee.Cluster.GetProxies() {
		result[proxy.Name] = proxy.GetPlayerlist()
	}

	response.SuccessWithData(c, gin.H{
		"sync_time": bungee.Cluster.LastFetch,
		"data":      result,
	})
}
