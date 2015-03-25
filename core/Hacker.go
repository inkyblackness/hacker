package core

import (
	"io/ioutil"
	"os"

	"github.com/inkyblackness/hacker/styling"
)

// Hacker is the main entry point for the hacker logic.
type Hacker struct {
	style styling.Style
}

// NewHacker returns a hacker instance to work with.
func NewHacker(style styling.Style) *Hacker {
	hacker := &Hacker{style: style}

	return hacker
}

// Load tries to load the data files from the two given directories. The second directory
// is optional.
func (hacker *Hacker) Load(path1, path2 string) string {
	files1, err1 := ioutil.ReadDir(path1)
	files2, err2 := ioutil.ReadDir(path2)
	result := ""

	fileNames := func(files []os.FileInfo) (names []string) {
		for _, file := range files {
			if !file.IsDir() {
				names = append(names, file.Name())
			}
		}
		return
	}

	if err1 == nil && err2 == nil {
		release := FindRelease(fileNames(files1), fileNames(files2))

		if release != nil {
			result = hacker.style.Status()("Determined Release [", release.name, "]")
		} else {
			result = hacker.style.Error()("Could not resolve release")
		}
	} else {
		result = hacker.style.Error()("Can't access directories")
	}

	return result
}
