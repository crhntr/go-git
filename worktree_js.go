package git

import (
	"github.com/go-git/go-git/v5/plumbing/format/index"
)

func isSymlinkWindowsNonAdmin(_ error) bool {
	return false
}


func init() {
	fillSystemInfo = func(_ *index.Entry, _ interface{}) {}
}