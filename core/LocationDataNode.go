package core

type locationDataNode struct {
	parentNode   DataNode
	dataLocation DataLocation
	filePath     string
	fileNames    []string

	fileDataNodeProvider FileDataNodeProvider
	fileDataNodes        map[string]DataNode
}

func newLocationDataNode(parentNode DataNode, dataLocation DataLocation,
	filePath string, fileNames []string, fileDataNodeProvider FileDataNodeProvider) *locationDataNode {
	node := &locationDataNode{
		parentNode:           parentNode,
		dataLocation:         dataLocation,
		filePath:             filePath,
		fileNames:            fileNames,
		fileDataNodeProvider: fileDataNodeProvider,
		fileDataNodes:        make(map[string]DataNode)}

	return node
}

func (node *locationDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *locationDataNode) Info() string {
	info := "Location: " + string(node.dataLocation) + "\n"
	info = info + "FilePath: [" + node.filePath + "]"

	return info
}

func (node *locationDataNode) ID() string {
	return string(node.dataLocation)
}

func (node *locationDataNode) Resolve(path string) (resolved DataNode) {
	temp, existing := node.fileDataNodes[path]

	if existing {
		resolved = temp
	} else if node.isFileKnown(path) {
		resolved = node.fileDataNodeProvider.Provide(node, node.filePath, path)
		if resolved != nil {
			node.fileDataNodes[path] = resolved
		}
	}

	return
}

func (node *locationDataNode) isFileKnown(fileName string) (result bool) {
	for _, knownName := range node.fileNames {
		if knownName == fileName {
			result = true
		}
	}
	return
}
