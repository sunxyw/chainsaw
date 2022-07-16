package requests

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type McmapRequest struct {
	Name string                `json:"name,omitempty" valid:"name"`
	Type string                `json:"type,omitempty" valid:"type"`
	File *multipart.FileHeader `valid:"file" form:"file"`
}

func McmapSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":      []string{"required"},
		"type":      []string{"required"},
		"file:file": []string{"required", "ext:zip,7z,schematic", "size:20971520" /* 20MB */},
	}
	messages := govalidator.MapData{
		"name": []string{"required:名称不能为空"},
		"type": []string{"required:类型不能为空"},
		"file:file": []string{
			"required:文件不能为空",
			"ext:文件类型必须是 zip,7z,schematic",
			"size:文件大小必须小于 20MB",
		},
	}
	return validateFile(c, data, rules, messages)
}
