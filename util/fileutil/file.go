package fileutil

import (
	"os"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return true
	}
}
