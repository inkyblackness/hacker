package core

import (
	"bytes"
	"path/filepath"
	"strings"

	chunkDos "github.com/inkyblackness/res/chunk/dos"
	textDos "github.com/inkyblackness/res/textprop/dos"
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
		reader := bytes.NewReader(rawData)

		if lowercaseFileName == "objprop.dat" {
		} else if lowercaseFileName == "textprop.dat" {
			propProvider, propErr := textDos.NewProvider(reader)

			if propErr == nil {
				node = NewTexturePropertiesDataNode(parentNode, fileName, propProvider)
			}
		} else {
			chunkProvider, chunkErr := chunkDos.NewChunkProvider(reader)

			if chunkErr == nil {
				node = NewResourceDataNode(parentNode, fileName, chunkProvider)
			}
		}
	}

	return
}
