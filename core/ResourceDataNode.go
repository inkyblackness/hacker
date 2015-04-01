package core

import (
	"github.com/inkyblackness/res/chunk"
)

type resourceDataNode struct {
	parentNode DataNode
	name       string

	provider       chunk.Provider
	chunkDataNodes []*chunkDataNode
}

func NewResourceDataNode(parentNode DataNode, name string, provider chunk.Provider) DataNode {
	ids := provider.IDs()
	node := &resourceDataNode{parentNode: parentNode,
		name:           name,
		chunkDataNodes: make([]*chunkDataNode, len(ids))}
	for index, id := range ids {
		node.chunkDataNodes[index] = newChunkDataNode(node, id)
	}

	return node
}

func (node *resourceDataNode) Parent() DataNode {
	return node.parentNode
}

func (node *resourceDataNode) Info() string {
	info := "ResourceFile: " + node.name + "\n"
	info += "IDs:"
	for _, node := range node.chunkDataNodes {
		info += " " + node.ID()
	}

	return info
}

func (node *resourceDataNode) ID() string {
	return node.name
}

func (node *resourceDataNode) Resolve(path string) DataNode {
	return nil
}
