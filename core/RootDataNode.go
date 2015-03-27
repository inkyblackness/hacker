package core

type rootDataNode struct {
	release   *ReleaseDesc
	locations map[DataLocation]*locationDataNode
}

func newRootDataNode(release *ReleaseDesc, hdLocation, cdLocation *locationDataNode) *rootDataNode {
	node := &rootDataNode{
		release:   release,
		locations: map[DataLocation]*locationDataNode{HD: hdLocation}}

	if cdLocation != nil {
		node.locations[CD] = cdLocation
	}

	return node
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
