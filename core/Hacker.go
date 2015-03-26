package core

import (
	"github.com/inkyblackness/hacker/styling"
)

// Hacker is the main entry point for the hacker logic.
type Hacker struct {
	style      styling.Style
	fileAccess *fileAccess

	root *rootDataNode
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
	var release *ReleaseDesc
	result := ""
	var hdLocation *locationDataNode
	var cdLocation *locationDataNode

	if err1 != nil {
		result = hacker.style.Error()("Can't access directories")
	} else if len(path2) == 0 {
		fileNames1 := fileNames(files1)
		release = FindRelease(fileNames1, nil)
		hdLocation = newLocationDataNode(HD, path1, fileNames1)
	} else {
		files2, err2 := hacker.fileAccess.readDir(path2)

		if err2 == nil {
			fileNames1 := fileNames(files1)
			fileNames2 := fileNames(files2)

			release = FindRelease(fileNames1, fileNames2)
			if release == nil {
				release = FindRelease(fileNames2, fileNames1)
				hdLocation = newLocationDataNode(HD, path2, fileNames2)
				cdLocation = newLocationDataNode(CD, path1, fileNames1)
			} else {
				hdLocation = newLocationDataNode(HD, path1, fileNames1)
				cdLocation = newLocationDataNode(CD, path2, fileNames2)
			}
		} else {
			result = hacker.style.Error()("Can't access directories")
		}
	}
	if release != nil {
		hacker.root = newRootDataNode(release, hdLocation, cdLocation)
		result = hacker.style.Status()("Loaded release [", release.name, "]")
	} else if len(result) == 0 {
		result = hacker.style.Error()("Could not resolve release")
	}

	return result
}

// Info returns the status of the current node
func (hacker *Hacker) Info() string {
	return hacker.style.Error()("not implemented")
}
