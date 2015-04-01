package core

import (
	"strings"

	"github.com/inkyblackness/hacker/styling"
)

// Hacker is the main entry point for the hacker logic.
type Hacker struct {
	style      styling.Style
	fileAccess *fileAccess

	root    *rootDataNode
	curNode dataNode
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
	var root *rootDataNode
	result := ""

	if err1 != nil {
		result = hacker.style.Error()("Can't access directories")
	} else if len(path2) == 0 {
		fileNames1 := fileNames(files1)
		release = FindRelease(fileNames1, nil)
		root = newRootDataNode(release)
		root.addLocation(newLocationDataNode(root, HD, path1, fileNames1))
	} else {
		files2, err2 := hacker.fileAccess.readDir(path2)

		if err2 == nil {
			fileNames1 := fileNames(files1)
			fileNames2 := fileNames(files2)

			release = FindRelease(fileNames1, fileNames2)
			if release == nil {
				release = FindRelease(fileNames2, fileNames1)
				root = newRootDataNode(release)
				root.addLocation(newLocationDataNode(root, HD, path2, fileNames2))
				root.addLocation(newLocationDataNode(root, CD, path1, fileNames1))
			} else {
				root = newRootDataNode(release)
				root.addLocation(newLocationDataNode(root, HD, path1, fileNames1))
				root.addLocation(newLocationDataNode(root, CD, path2, fileNames2))
			}
		} else {
			result = hacker.style.Error()("Can't access directories")
		}
	}
	if release != nil {
		hacker.root = root
		hacker.curNode = root
		result = hacker.style.Status()("Loaded release [", release.name, "]")
	} else if len(result) == 0 {
		result = hacker.style.Error()("Could not resolve release")
	}

	return result
}

// Info returns the status of the current node
func (hacker *Hacker) Info() string {
	var result string

	if hacker.curNode != nil {
		result = hacker.curNode.info()
	} else {
		result = hacker.style.Error()(`No data loaded. Use the [load "path1" "path2"] command.`)
	}

	return result
}

func (hacker *Hacker) ChangeDirectory(path string) (result string) {
	parts := strings.Split(path, "/")
	tempNode := hacker.curNode

	if parts[0] == "" {
		tempNode = hacker.root
	}
	for _, part := range parts {
		if tempNode != nil && part != "" {
			if part == ".." {
				tempNode = tempNode.parent()
			} else {
				tempNode = tempNode.resolve(part)
			}
		}
	}
	if tempNode != nil {
		hacker.curNode = tempNode
		result = ""
	} else {
		result = hacker.style.Error()(`Directory not found: ""`, path)
	}
	return
}
