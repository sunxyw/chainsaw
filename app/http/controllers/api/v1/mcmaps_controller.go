package v1

import (
	"gohub/app/models/mcmap"
	"gohub/app/requests"
	"gohub/pkg/file"
	"gohub/pkg/logger"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type McmapsController struct {
	BaseAPIController
}

func (ctrl *McmapsController) Index(c *gin.Context) {
	mcmaps, paging := mcmap.Paginate(c, 10)
	response.SuccessWithData(c, gin.H{
		"data":   mcmaps,
		"paging": paging,
	})
}

func (ctrl *McmapsController) Show(c *gin.Context) {
	mcmapModel := mcmap.Get(c.Param("id"))
	if mcmapModel.ID == 0 {
		response.NotFound(c)
		return
	}
	response.SuccessWithData(c, gin.H{
		"data": mcmapModel,
	})
}

func (ctrl *McmapsController) Store(c *gin.Context) {

	request := requests.McmapRequest{}
	if ok := requests.Validate(c, &request, requests.McmapSave); !ok {
		return
	}

	mapFile, err := file.SaveUploadFile(c, request.File, "mcmap")
	if err != nil {
		response.ServerError(c, "保存文件失败")
		return
	}

	mcmapModel := mcmap.Mcmap{
		Name:    request.Name,
		Type:    request.Type,
		Address: "file://" + mapFile,
		Size:    request.File.Size,
	}
	logger.Dump(mcmapModel)
	return
	mcmapModel.Create()
	if mcmapModel.ID > 0 {
		response.Created(c, gin.H{
			"data": mcmapModel,
		})
	} else {
		response.ServerError(c, "创建失败，请稍后尝试~")
	}
}
