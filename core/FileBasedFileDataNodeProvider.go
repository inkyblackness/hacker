package core

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/inkyblackness/res/chunk/dos"
)

type fileBasedFileDataNodeProvider struct {
	access fileAccess
}

func newFileDataNodeProvider(access fileAccess) FileDataNodeProvider {
	provider := &fileBasedFileDataNodeProvider{
		access: access}

	return provider
}

// Provide tries to resolve and return a DataNode for the given file.
func (provider *fileBasedFileDataNodeProvider) Provide(parentNode DataNode, filePath, fileName string) (node DataNode) {
	rawData, err := provider.access.readFile(filepath.Join(filePath, fileName))

	if err == nil {
		lowercaseFileName := strings.ToLower(fileName)

		if lowercaseFileName == "objprop.dat" {
		} else if lowercaseFileName == "textprop.dat" {
		} else {
			reader := bytes.NewReader(rawData)
			chunkProvider, chunkErr := dos.NewChunkProvider(reader)

			if chunkErr == nil {
				node = NewResourceDataNode(parentNode, fileName, chunkProvider)
			}
		}
	}

	return
}
