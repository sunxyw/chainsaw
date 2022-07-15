package v1

import (
	"gohub/app/models/news"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type NewsController struct {
	BaseAPIController
}

func (ctrl *NewsController) Index(c *gin.Context) {
	news, paging := news.Paginate(c, 10)
	response.SuccessWithData(c, gin.H{
		"data":   news,
		"paging": paging,
	})
}
