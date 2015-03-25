package core

import (
	"os"
)

func fileNames(files []os.FileInfo) (names []string) {
	for _, file := range files {
		if !file.IsDir() {
			names = append(names, file.Name())
		}
	}
	return
}
