package core

type locationDataNode struct {
	parentNode   DataNode
	dataLocation DataLocation
	filePath     string
}

func newLocationDataNode(parentNode DataNode, dataLocation DataLocation, filePath string, files []string) *locationDataNode {
	node := &locationDataNode{parentNode: parentNode, dataLocation: dataLocation, filePath: filePath}

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

func (node *locationDataNode) Resolve(path string) DataNode {
	return nil
}
