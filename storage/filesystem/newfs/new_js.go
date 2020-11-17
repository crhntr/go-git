// +build js

package newfs

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
)

func New(_ string) billy.Filesystem {
	return memfs.New()
}