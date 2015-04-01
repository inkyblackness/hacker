package core

import (
	"fmt"
	"strings"

	"github.com/inkyblackness/hacker/styling"
)

// Hacker is the main entry point for the hacker logic.
type Hacker struct {
	style                styling.Style
	fileAccess           fileAccess
	fileDataNodeProvider FileDataNodeProvider

	root    *rootDataNode
	curNode DataNode
}

// NewHacker returns a hacker instance to work with.
func NewHacker(style styling.Style) *Hacker {
	access := realFileAccess
	hacker := &Hacker{
		style:                style,
		fileAccess:           access,
		fileDataNodeProvider: newFileDataNodeProvider(access)}

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
		root.addLocation(newLocationDataNode(root, HD, path1, fileNames1, hacker.fileDataNodeProvider))
	} else {
		files2, err2 := hacker.fileAccess.readDir(path2)

		if err2 == nil {
			fileNames1 := fileNames(files1)
			fileNames2 := fileNames(files2)

			release = FindRelease(fileNames1, fileNames2)
			if release == nil {
				release = FindRelease(fileNames2, fileNames1)
				root = newRootDataNode(release)
				root.addLocation(newLocationDataNode(root, HD, path2, fileNames2, hacker.fileDataNodeProvider))
				root.addLocation(newLocationDataNode(root, CD, path1, fileNames1, hacker.fileDataNodeProvider))
			} else {
				root = newRootDataNode(release)
				root.addLocation(newLocationDataNode(root, HD, path1, fileNames1, hacker.fileDataNodeProvider))
				root.addLocation(newLocationDataNode(root, CD, path2, fileNames2, hacker.fileDataNodeProvider))
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
		result = hacker.curNode.Info()
	} else {
		result = hacker.style.Error()(`No data loaded. Use the [load "path1" "path2"] command.`)
	}

	return result
}

// CurrentDirectory returns the absolute path to the current directory in string form
func (hacker *Hacker) CurrentDirectory() string {
	tempNode := hacker.curNode
	path := ""

	for tempNode != nil && tempNode != hacker.root {
		path = "/" + tempNode.ID() + path
		tempNode = tempNode.Parent()
	}

	return path
}

// ChangeDirectory changes the currently active node
func (hacker *Hacker) ChangeDirectory(path string) (result string) {
	parts := strings.Split(path, "/")
	tempNode := hacker.curNode

	if parts[0] == "" {
		tempNode = hacker.root
	}
	for _, part := range parts {
		if tempNode != nil && part != "" {
			if part == ".." {
				tempNode = tempNode.Parent()
			} else {
				tempNode = tempNode.Resolve(part)
			}
		}
	}
	if tempNode != nil {
		hacker.curNode = tempNode
		result = ""
	} else {
		result = hacker.style.Error()(`Directory not found: "`, path, `"`)
	}
	return
}

func (hacker *Hacker) Dump() (result string) {
	rightPad := func(input string, length int) string {
		padded := fmt.Sprintf(fmt.Sprintf("%%s%%%ds", length), input, "")
		return padded[0:length]
	}

	if hacker.curNode != nil {
		data := hacker.curNode.Data()
		hexDump := ""
		asciiDump := ""

		addLine := func(offset int) {
			result = result + fmt.Sprintf("%04X %s  %s\n", offset, rightPad(hexDump, 49), rightPad(asciiDump, 17))
			hexDump = ""
			asciiDump = ""
		}

		for index, value := range data {
			if index == 0 {
			} else if (index % 16) == 0 {
				addLine(((index / 16) - 1) * 16)
			} else if (index % 8) == 0 {
				hexDump += " "
				asciiDump += " "
			}

			hexDump += fmt.Sprintf(" %02X", value)
			if value >= 0x20 && value < 0x80 {
				asciiDump += string(value)
			} else {
				asciiDump += "."
			}
		}
		if hexDump != "" {
			addLine((len(data) / 16) * 16)
		}
	}
	return
}
