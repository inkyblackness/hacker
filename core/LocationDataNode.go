package core

type locationDataNode struct {
	parentDataNode

	dataLocation DataLocation
	filePath     string
	fileNames    []string
}

func newLocationDataNode(parentNode DataNode, dataLocation DataLocation,
	filePath string, fileNames []string, fileDataNodeProvider FileDataNodeProvider) *locationDataNode {
	node := &locationDataNode{
		parentDataNode: makeParentDataNode(parentNode, dataLocation.String(), len(fileNames)),
		dataLocation:   dataLocation,
		filePath:       filePath,
		fileNames:      fileNames}

	node.setChildResolver(func(path string) (resolved DataNode) {
		if node.isFileKnown(path) {
			resolved = fileDataNodeProvider.Provide(node, node.filePath, path)
		}
		return
	})

	return node
}

func (node *locationDataNode) Info() string {
	info := "Location: " + string(node.dataLocation) + "\n"
	info = info + "FilePath: [" + node.filePath + "]\n"
	info = info + "Files:"
	for _, fileName := range node.fileNames {
		info = info + " " + fileName
	}

	return info
}

func (node *locationDataNode) isFileKnown(fileName string) (result bool) {
	for _, knownName := range node.fileNames {
		if knownName == fileName {
			result = true
		}
	}
	return
}
