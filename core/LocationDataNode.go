package core

type locationDataNode struct {
	dataLocation DataLocation
	filePath     string
}

func newLocationDataNode(dataLocation DataLocation, filePath string, files []string) *locationDataNode {
	node := &locationDataNode{dataLocation: dataLocation, filePath: filePath}

	return node
}
