// Package file 文件操作辅助函数
package file

import (
	"fmt"
	"gohub/pkg/app"
	"gohub/pkg/auth"
	"gohub/pkg/helpers"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

var (
	filesystems map[string]afero.Fs
	fsHelpers   map[string]afero.Afero
)

// AddFileSystem 添加文件系统
func AddFileSystem(name string, fs afero.Fs) {
	filesystems[name] = fs
}

// Fs 获取文件系统
func Fs(name string) afero.Fs {
	return filesystems[name]
}

func FsHelper(name string) afero.Afero {
	if _, ok := fsHelpers[name]; !ok {
		fsHelpers[name] = afero.Afero{Fs: Fs(name)}
	}
	return fsHelpers[name]
}

// Put 将数据存入文件
func Put(data []byte, to string) error {
	return FsHelper("local").WriteFile(to, data, 0644)
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	exists, err := FsHelper("local").Exists(fileToCheck)
	if err != nil {
		return false
	}
	return exists
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {

	fs := Fs("public")

	var avatar string
	// 确保目录存在，不存在创建
	dirName := fmt.Sprintf("avatars/%s/%s/", app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	fs.MkdirAll(dirName, 0755)

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/avatars/2021/12/22/1/nFDacgaWKpWWOmOt.png
	avatarPath := dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	// 裁切图片
	img, err := imaging.Open("public/uploads"+avatarPath, imaging.AutoOrientation(true))
	if err != nil {
		return avatar, err
	}
	resizeAvatar := imaging.Thumbnail(img, 256, 256, imaging.Lanczos)
	resizeAvatarName := randomNameFromUploadFile(file)
	resizeAvatarPath := dirName + resizeAvatarName
	err = imaging.Save(resizeAvatar, resizeAvatarPath)
	if err != nil {
		return avatar, err
	}

	// 删除老文件
	err = fs.Remove(avatarPath)
	if err != nil {
		return avatar, err
	}

	return dirName + resizeAvatarName, nil
}

func SaveUploadFile(c *gin.Context, file *multipart.FileHeader, folder string) (string, error) {

	var fileName string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/%s/%s/", folder, app.TimenowInTimezone().Format("2006/01/02"))
	Fs("local").MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName = randomNameFromUploadFile(file)
	// public/uploads/files/2021/12/22/1/nFDacgaWKpWWOmOt.png
	filePath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return fileName, err
	}

	return dirName + fileName, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
