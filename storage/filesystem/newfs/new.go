// +build !js

package newfs

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

func New(baseDir string) billy.Filesystem {
	return osfs.New(baseDir)
}