package core

type locationDataNode struct {
	parentNode   dataNode
	dataLocation DataLocation
	filePath     string
}

func newLocationDataNode(parentNode dataNode, dataLocation DataLocation, filePath string, files []string) *locationDataNode {
	node := &locationDataNode{parentNode: parentNode, dataLocation: dataLocation, filePath: filePath}

	return node
}

func (node *locationDataNode) parent() dataNode {
	return node.parentNode
}

func (node *locationDataNode) info() string {
	info := "FilePath: [" + node.filePath + "]"

	return info
}

func (node *locationDataNode) resolve(path string) dataNode {
	return nil
}
