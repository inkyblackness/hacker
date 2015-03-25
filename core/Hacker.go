package core

import (
	"github.com/inkyblackness/hacker/styling"
)

// Hacker is the main entry point for the hacker logic.
type Hacker struct {
	style      styling.Style
	fileAccess *fileAccess
}

// NewHacker returns a hacker instance to work with.
func NewHacker(style styling.Style) *Hacker {
	hacker := &Hacker{style: style, fileAccess: &realFileAccess}

	return hacker
}

// Load tries to load the data files from the two given directories. The second directory
// is optional.
func (hacker *Hacker) Load(path1, path2 string) string {
	files1, err1 := hacker.fileAccess.readDir(path1)
	result := ""

	tryLoad := func(hdFileNames, cdFileNames []string) (loaded bool) {
		release := FindRelease(hdFileNames, cdFileNames)

		if release != nil {
			result = hacker.style.Status()("Loaded release [", release.name, "]")
			loaded = true
		} else {
			result = hacker.style.Error()("Could not resolve release")
		}
		return
	}

	if err1 != nil {
		result = hacker.style.Error()("Can't access directories")
	} else if len(path2) == 0 {
		tryLoad(fileNames(files1), nil)
	} else {
		files2, err2 := hacker.fileAccess.readDir(path2)

		if err2 == nil {
			fileNames1 := fileNames(files1)
			fileNames2 := fileNames(files2)

			if !tryLoad(fileNames1, fileNames2) {
				tryLoad(fileNames2, fileNames1)
			}
		} else {
			result = hacker.style.Error()("Can't access directories")
		}
	}

	return result
}
