package core

import (
	"github.com/inkyblackness/res/chunk"
)

type resourceDataNode struct {
	parentDataNode
}

func NewResourceDataNode(parentNode DataNode, name string, provider chunk.Provider) DataNode {
	ids := provider.IDs()
	node := &resourceDataNode{parentDataNode: makeParentDataNode(parentNode, name, len(ids))}

	for _, id := range ids {
		node.addChild(newChunkDataNode(node, id, provider.Provide(id)))
	}

	return node
}

func (node *resourceDataNode) Info() string {
	info := "ResourceFile: " + node.ID() + "\n"
	info += "IDs:"
	for _, node := range node.Children() {
		info += " " + node.ID()
	}

	return info
}
