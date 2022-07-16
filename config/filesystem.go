// Package config 站点配置信息
package config

import (
	"gohub/pkg/config"
	"gohub/pkg/file"

	"github.com/spf13/afero"
)

func init() {
	config.Add("filesystem", func() map[string]interface{} {
		return map[string]interface{}{

			"filesystems": []file.FsConf{
				{
					Name: "local",
					Fs:   afero.NewBasePathFs(afero.NewOsFs(), "storage"),
				},
				{
					Name: "public",
					Fs:   afero.NewBasePathFs(afero.NewOsFs(), "public/uploads"),
				},
				{
					Name: "memory",
					Fs:   afero.NewMemMapFs(),
				},
			},
		}
	})
}
