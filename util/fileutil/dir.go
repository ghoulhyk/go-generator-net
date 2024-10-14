package fileutil

import "os"

func IsDir(dirPath string) bool {
	stat, err := os.Stat(dirPath)
	if err != nil {
		return false
	} else {
		return stat.IsDir()
	}
}
