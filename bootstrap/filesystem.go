package bootstrap

import (
	"gohub/pkg/config"
	"gohub/pkg/file"
)

func SetupFilesystem() {
	for _, fsConf := range config.Get[[]file.FsConf]("filesystem.filesystems") {
		file.AddFileSystem(fsConf.Name, fsConf.Fs)
	}
}
