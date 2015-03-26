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
	return node.release.name
}
