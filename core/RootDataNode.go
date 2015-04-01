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

func (node *rootDataNode) Parent() DataNode {
	return nil
}

func (node *rootDataNode) Info() string {
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

func (node *rootDataNode) ID() string {
	return ""
}

func (node *rootDataNode) Resolve(path string) (resolved DataNode) {
	location, existing := node.locations[DataLocation(path)]

	if existing {
		resolved = location
	}
	return
}
