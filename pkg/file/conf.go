package file

import "github.com/spf13/afero"

type FsConf struct {
	Name string
	Fs   afero.Fs
}
