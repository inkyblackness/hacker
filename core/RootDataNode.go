package core

type rootDataNode struct {
	release   *ReleaseDesc
	locations map[DataLocation]*locationDataNode
}

func newRootDataNode(release *ReleaseDesc) *rootDataNode {
	node := &rootDataNode{
		release:   release,
		locations: make(map[DataLocation]*locationDataNode)}

	return node
}

func (node *rootDataNode) addLocation(location *locationDataNode) {
	node.locations[location.dataLocation] = location
}

func (node *rootDataNode) parent() dataNode {
	return nil
}

func (node *rootDataNode) info() string {
	info := "Release: [" + node.release.name + "]"
	info = info + "\nAvailable data locations:"
	if _, existing := node.locations[HD]; existing {
		info = info + " hd"
	}
	if _, existing := node.locations[CD]; existing {
		info = info + " cd"
	}

	return info
}

func (node *rootDataNode) resolve(path string) dataNode {
	return node.locations[DataLocation(path)]
}
